package tests

import (
	"os"
	"time"

	cpb "github.com/SailGame/GoDock/pb/core"
	ui "github.com/gizak/termui/v3"

	connmock "github.com/SailGame/GoDock/conn/mocks"
	"github.com/SailGame/GoDock/jui"
	juimock "github.com/SailGame/GoDock/jui/mocks"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func init(){
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)
	f, err := os.OpenFile("godock_test.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
}

type Fixture struct {
	mStore jui.Store
	mRouter *juimock.MockRouter
	mClient *connmock.MockGameCoreClient

	mPollUIEventC  chan ui.Event
	mCoreMsgEventC chan *cpb.BroadcastMsg
	mTimeTickC chan time.Time
}

func NewFixture(ctrl *gomock.Controller) *Fixture{
	fixture := &Fixture{
		mStore: jui.NewDefaultStore(),
		mRouter: juimock.NewMockRouter(ctrl),
		mClient: connmock.NewMockGameCoreClient(ctrl),
		mPollUIEventC: make(chan ui.Event),
		mCoreMsgEventC: make(chan *cpb.BroadcastMsg),
		mTimeTickC: make(chan time.Time),
	}
	return fixture
}