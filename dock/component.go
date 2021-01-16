package dock

import (
	ui "github.com/gizak/termui/v3"
	cpb "github.com/SailGame/GoDock/pb/core"
)

type Component interface {
	GetGrid() *ui.Grid
	HandleUIEvent(ui.Event) bool
	HandleServerEvent(*cpb.BroadcastMsg) bool
	// clear all internal state
	Reset() error
	// the standard timetick is called every 0.05 sec
	TimeTick()
}