package jui

import (
	ui "github.com/gizak/termui/v3"
	cpb "github.com/SailGame/GoDock/pb/core"
)

//go:generate mockgen -destination=mocks/component.go -package=mocks . Component
type Component interface {
	GetGrid() *ui.Grid
	HandleUIEvent(ui.Event) bool
	HandleServerEvent(*cpb.BroadcastMsg) bool
	// lifecycle
	WillMount(interface{})
	DidMount()
	WillUnmount()
	DidUnmount()
	Reset() error
	// the standard timetick is called every 0.05 sec
	TimeTick()
}