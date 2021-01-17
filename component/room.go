package component

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/SailGame/GoDock/component/data"
	cpb "github.com/SailGame/GoDock/pb/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Room struct {
	Default
	mRoomId int32
	mUsers []data.RoomUser
	mReady bool
	mExited bool
}

func NewRoom(d *Default) *Room{
	room := &Room{
		*d,
		0,
		make([]data.RoomUser, 0),
		false,
		false,
	}
	return room
}

func (room *Room) HandleUIEvent(e ui.Event) bool{
	log.Debugf("Room Recv UI event: %s", e.ID)
	switch e.ID {
	case "r":
		room.ready()
	case "s":
		room.set()
	default:
		return false
	}
	return true
}

func (room *Room) HandleServerEvent(*cpb.BroadcastMsg) bool{
	return false
}

// lifecycle
func (room *Room) WillMount(interface{}){
	room.refresh()
}

func (room *Room) WillUnmount(){
	log.Debugf("Room Will Unmount")

	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := room.mStore.GetGameCoreClient().ExitRoom(ctx, &cpb.ExitRoomArgs{Token: room.mStore.GetToken()})
	if err != nil {
		log.Fatalf("ExitRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		switch ret.Err {
		case cpb.ErrorNumber_ExitRoom_InvalidToken: // TODO: display this msg
		case cpb.ErrorNumber_ExitRoom_NotInRoom: // TODO: display this msg
		}
	}
}

func (room *Room) Reset() error{
	// TODO
	return nil
}

func (room *Room) ready(){
	log.Debugf(fmt.Sprintf("Room Change Ready State. Curr(%v)", room.mReady))

	var grpcReady cpb.Ready = cpb.Ready_UNSET
	if room.mReady{
		grpcReady = cpb.Ready_CANCEL
	}else{
		grpcReady = cpb.Ready_READY
	}
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := room.mStore.GetGameCoreClient().OperationInRoom(
		ctx,
		&cpb.OperationInRoomArgs{
			Token: room.mStore.GetToken(),
			RoomOperation: &cpb.OperationInRoomArgs_Ready{ Ready: grpcReady },
	})

	if err != nil {
		log.Fatalf("Ready %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		switch ret.Err {
		case cpb.ErrorNumber_OperRoom_CannotChangeReadyState: // TODO
		return
		}
	}
	room.mReady = !room.mReady
	room.refresh()
}

func (room *Room) set(){

}

func (room *Room) refresh(){
	room.mGrid.Set(ui.NewRow(
		0.2/2, ui.NewCol(1, widgets.NewParagraph()),
	))
}