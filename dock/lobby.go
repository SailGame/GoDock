package dock

import (
	ui "github.com/gizak/termui/v3"
	cpb "github.com/SailGame/GoDock/pb/core"
)

type LobbyComponent struct {
	mGameCoreClient cpb.GameCoreClient
	mGrid *ui.Grid
	mEventHook map[string]func()
}

func NewLobbyComponent(gameCoreClient cpb.GameCoreClient) *LobbyComponent{
	lc := &LobbyComponent{
		mGameCoreClient: gameCoreClient,
		mGrid: ui.NewGrid(),
	}
	lc.mEventHook = map[string]func(){
		"<c>": lc.createRoom,
		"<l>": lc.listRoom,
	}
	return lc
}

func (lc *LobbyComponent) GetGrid() *ui.Grid{
	return lc.mGrid
}

func (lc *LobbyComponent) HandleUIEvent(e ui.Event) bool{
	hook, ok := lc.mEventHook[e.ID]
	if ok {
		hook()
		return true
	}
	return false
}

func (lc *LobbyComponent) HandleServerEvent(*cpb.BroadcastMsg) bool{
	return false
}

func (lc *LobbyComponent) Reset() error{
	return nil
}

func (lc *LobbyComponent) TimeTick(){

}

func (lc *LobbyComponent) createRoom(){

}

func (lc *LobbyComponent) listRoom(){

}