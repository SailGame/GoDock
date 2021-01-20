package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/SailGame/GoDock/component"
	"github.com/SailGame/GoDock/conn"
	"github.com/SailGame/GoDock/dock"
	"github.com/SailGame/GoDock/jui"
	ui "github.com/gizak/termui/v3"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:8080", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
	logFile            = flag.String("log_file", "./godock.log", "The log file of GoDock")
	logLevel           = flag.String("log_level", "Info", "The log level of GoDock (Debug/Info/Warn/Error)")
	userName           = flag.String("user_name", "sail", "Your name")
)

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})

	if *logLevel == "Debug" {
		log.SetLevel(log.DebugLevel)
	} else if *logLevel == "Info" {
		log.SetLevel(log.InfoLevel)
	} else if *logLevel == "Warn" {
		log.SetLevel(log.WarnLevel)
	} else if *logLevel == "Error" {
		log.SetLevel(log.ErrorLevel)
	} else {
		panic(fmt.Sprintf("Unknown logLevel: %s", *logLevel))
	}

	f, err := os.OpenFile(*logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

}

func initGrpcDialOption() []grpc.DialOption {
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			log.Fatal("empty caFile")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	return append(opts, grpc.WithBlock(), grpc.WithTimeout(time.Second * 3))
}

func initGrpcClient(grpcConn *grpc.ClientConn) *conn.GameCoreConn{
	// init grpc client
	coreClientConn := conn.NewGameCoreConn(grpcConn)
	err := coreClientConn.Login(*userName)
	if err != nil{
		log.Fatal(err)
	}
	err = coreClientConn.ListenToCore()
	if err != nil{
		log.Fatal(err)
	}

	go coreClientConn.LoopListenStream(func() {
		log.Fatalf("Disconnect from Core Server");
	})
	return coreClientConn
}

func init() {
	flag.Parse()
	initLog()
}

func main() {
	// init coreClientConn
	grpcConn, err := grpc.Dial(*serverAddr, initGrpcDialOption()...)
	if err != nil {
		log.Fatalf("failed to connect to Core Server")
	}
	defer grpcConn.Close()
	coreClientConn := initGrpcClient(grpcConn)

	// init ui
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// init jui
	store := jui.NewDefaultStore()
	router := jui.NewDefaultRouter()
	store.SetRouter(router)
	store.SetToken(coreClientConn.GetToken())
	store.SetGameCoreClient(coreClientConn.GetGameCoreClient())

	// init components
	component.Init(store)
	router.Navigate("/", nil)
	// init dock
	ticker := time.NewTicker(time.Millisecond * 100).C
	gameDock := dock.NewDock(store, ui.PollEvents(), coreClientConn.GetBroadcastMsgCh(), ticker)

	log.Info("Game Dock Event Loop start")
	gameDock.Loop()
}