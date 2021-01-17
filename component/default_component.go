package component

import (
	"github.com/SailGame/GoDock/jui"
	cpb "github.com/SailGame/GoDock/pb/core"
	ui "github.com/gizak/termui/v3"
)


type Default struct {
	mStore jui.Store
	mGrid *ui.Grid
}

func NewDefaultComponent(store jui.Store) *Default{
	return &Default{
		mStore: store,
		mGrid: ui.NewGrid(),
	}
}

func (d *Default) GetGrid() *ui.Grid {
	return d.mGrid
}

func (d *Default) HandleUIEvent(ui.Event) bool{
	return false
}

func (d *Default) HandleServerEvent(*cpb.BroadcastMsg) bool{
	return false
}

// lifecycle
func (d *Default) WillMount(interface{}){

}
func (d *Default) DidMount(){

}
func (d *Default) WillUnmount(){

}
func (d *Default) DidUnmount(){

}
func (d *Default) Reset() error{
	return nil
}
// the standard timetick is called every 0.05 sec
func (d *Default) TimeTick(){

}