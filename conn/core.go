package conn

import (
	"context"
	"time"

	cpb "github.com/SailGame/GoDock/pb/core"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GameCoreConn struct {
	mCtx context.Context
	mCoreClient cpb.GameCoreClient
	mLisClient cpb.GameCore_ListenClient
	mMsgC	chan *cpb.BroadcastMsg
	mToken string
}

func NewGameCoreConn(ctx context.Context, grpcConn *grpc.ClientConn) *GameCoreConn {
	return &GameCoreConn{
		mCtx: ctx,
		mCoreClient: cpb.NewGameCoreClient(grpcConn),
		mMsgC: make(chan *cpb.BroadcastMsg),
	}
}

func (gcc *GameCoreConn) ListenToCore() error {
	var err error
	gcc.mLisClient, err = gcc.mCoreClient.Listen(gcc.mCtx, &cpb.ListenArgs{
		Token: gcc.mToken,
	})
	if err != nil {
		return err
	}
	return err
}

func (gcc *GameCoreConn) LoopListenStream(onStop func()){
	for{
		msg, err := gcc.mLisClient.Recv()
		if err != nil {
			log.Warn(err)
			onStop()
			return
		}
		log.Debugf("GameCoreConn Recv Msg %v", msg)
		gcc.mMsgC <- msg
	}
}

func (gcc *GameCoreConn) GetBroadcastMsgCh() <-chan *cpb.BroadcastMsg {
	return gcc.mMsgC
}

// GameCoreClient

// client
func (gcc *GameCoreConn) Login(in *cpb.LoginArgs) (*cpb.LoginRet){
	loginRet, err := gcc.mCoreClient.Login(gcc.mCtx, in)
	if err != nil{
		log.Fatal(err)
	}
	if loginRet.Err != cpb.ErrorNumber_OK{
		log.Fatal(loginRet.Err)
	}
	gcc.mToken = loginRet.Token
	return loginRet
}

func (gcc *GameCoreConn) QueryAccount(in *cpb.QueryAccountArgs) (*cpb.QueryAccountRet){
	return nil
}
// room
func (gcc *GameCoreConn) CreateRoom(in *cpb.CreateRoomArgs) (*cpb.CreateRoomRet){
	ctx, _ := context.WithTimeout(gcc.mCtx, 3 * time.Second)
	in.Token = gcc.mToken
	ret, err := gcc.mCoreClient.CreateRoom(ctx, in)
	if err != nil {
		log.Fatalf("CreateRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		log.Fatalf("CreateRoom %v", ret.Err)
	}
	return ret
}

func (gcc *GameCoreConn) ControlRoom(in *cpb.ControlRoomArgs) (*cpb.ControlRoomRet){
	return nil
}

func (gcc *GameCoreConn) ListRoom(in *cpb.ListRoomArgs) (*cpb.ListRoomRet){
	ctx, _ := context.WithTimeout(gcc.mCtx, 3 * time.Second)
	ret, err := gcc.mCoreClient.ListRoom(ctx, in)
	if err != nil {
		log.Fatalf("ListRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		log.Fatalf("ListRoom %v", ret.Err)
	}
	return ret
}

func (gcc *GameCoreConn) JoinRoom(in *cpb.JoinRoomArgs) (*cpb.JoinRoomRet){
	ctx, _ := context.WithTimeout(gcc.mCtx, 3 * time.Second)
	in.Token = gcc.mToken
	ret, err := gcc.mCoreClient.JoinRoom(ctx, in)
	if err != nil {
		log.Fatalf("JoinRoom %v", err)
	}
	return ret
}
func (gcc *GameCoreConn) ExitRoom(in *cpb.ExitRoomArgs) (*cpb.ExitRoomRet){
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	in.Token = gcc.mToken
	ret, err := gcc.mCoreClient.ExitRoom(ctx, in)
	if err != nil {
		log.Fatalf("ExitRoom %v", err)
	}
	return ret
}
func (gcc *GameCoreConn) QueryRoom(in *cpb.QueryRoomArgs) (*cpb.QueryRoomRet){
	return nil
}
// OperationInRoom
func (gcc *GameCoreConn) Ready(in *cpb.OperationInRoomArgs_Ready) (*cpb.OperationInRoomRet){
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := gcc.mCoreClient.OperationInRoom(
		ctx,
		&cpb.OperationInRoomArgs{
			Token: gcc.mToken,
			RoomOperation: in,
	})
	if err != nil {
		log.Fatalf("Ready %v", err)
	}
	return ret
}