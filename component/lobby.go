package component

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"

	"time"
	"github.com/SailGame/GoDock/component/data"
	cpb "github.com/SailGame/GoDock/pb/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Lobby struct {
	Default
	mRooms []data.Room
	mListGameName string
	mSelectedRoom int
}

func NewLobby(d *Default) *Lobby{
	lc := &Lobby{
		*d,
		make([]data.Room, 0),
		"",
		0,
	}
	return lc
}

func (lc *Lobby) HandleUIEvent(e ui.Event) bool{
	log.Debugf("Lobby Recv UI event: %s", e.ID)
	switch e.ID {
	case "l":
		lc.listRoom()
	case "c":
		lc.createRoom()
	case "j":
		lc.joinRoom()
	default:
		return false
	}
	return true
}

func (lc *Lobby) HandleServerEvent(*cpb.BroadcastMsg) bool{
	return false
}

// lifecycle
func (lc *Lobby) WillMount(interface{}){
	if len(lc.mRooms) == 0{
		lc.listRoom()
	}
}

func (lc *Lobby) Reset() error{
	lc.mRooms = lc.mRooms[:0]
	lc.mSelectedRoom = 0
	return nil
}

func (lc *Lobby) createRoom(){
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := lc.mStore.GetGameCoreClient().CreateRoom(ctx, &cpb.CreateRoomArgs{Token: lc.mStore.GetToken()})
	if err != nil {
		log.Fatalf("CreateRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		log.Fatalf("CreateRoom %v", ret.Err)
	}
	lc.mSelectedRoom = int(ret.RoomId)
	lc.joinRoom()
}

func (lc *Lobby) listRoom(){
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := lc.mStore.GetGameCoreClient().ListRoom(ctx, &cpb.ListRoomArgs{GameName: lc.mListGameName})
	if err != nil {
		log.Fatalf("ListRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		log.Fatalf("ListRoom %v", ret.Err)
	}
	lc.mRooms = lc.mRooms[:0]
	for _, v := range(ret.Room) {
		lc.mRooms = append(lc.mRooms, data.Room{*v})
	}
	lc.mSelectedRoom = 0
	lc.refresh()
}

func (lc *Lobby) joinRoom(){
	ctx, _ := context.WithTimeout(context.TODO(), 3 * time.Second)
	ret, err := lc.mStore.GetGameCoreClient().JoinRoom(ctx, &cpb.JoinRoomArgs{Token: lc.mStore.GetToken(), RoomId: int32(lc.mSelectedRoom)})
	if err != nil {
		log.Fatalf("JoinRoom %v", err)
	}
	if ret.Err != cpb.ErrorNumber_OK {
		switch ret.Err {
		case cpb.ErrorNumber_JoinRoom_FullRoom: // TODO: display this msg
		case cpb.ErrorNumber_JoinRoom_InvalidRoomID: // TODO: display this msg
		case cpb.ErrorNumber_JoinRoom_UserIsInAnotherRoom: // TODO: display this msg
			log.Fatalf("JoinRoom %v", ret.Err)
		case cpb.ErrorNumber_JoinRoom_InvalidToken:
			log.Fatalf("JoinRoom %v", ret.Err)
		}
		return
	}
	lc.mStore.GetRouter().Navigate(ROOM, nil)
}

func (lc *Lobby) refresh(){
	roomNum := len(lc.mRooms)
	roomRows := make([]interface{}, 0, roomNum)
	for i, r := range lc.mRooms{
		card := widgets.NewParagraph()
		if i == lc.mSelectedRoom {
			card.BorderStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
		}
		card.Text = fmt.Sprintf("Id: %d Game: %s PlayerNum: %d", r.GetRoomId(), r.GetGameName(), len(r.GetUserName()))
		roomRows = append(roomRows, ui.NewRow(1.0/float64(roomNum), ui.NewCol(1, card)))
	}
	lc.mGrid.Set(
		roomRows...
	)
}