package tests

import (
	"testing"

	"github.com/SailGame/GoDock/dock"
	cpb "github.com/SailGame/GoDock/pb/core"
	"github.com/golang/mock/gomock"
	// log "github.com/sirupsen/logrus"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
  
	fixture := NewFixture(ctrl)

	fixture.mClient.
	  EXPECT().
	  ListRoom(gomock.Any()).
	  Return(cpb.ListRoomRet{
		  Room: make([]*cpb.Room, 0),
	  }, nil)

	d := dock.NewDock(fixture.mStore, fixture.mPollUIEventC, fixture.mCoreMsgEventC, fixture.mTimeTickC)

	go d.Loop()
	
  }