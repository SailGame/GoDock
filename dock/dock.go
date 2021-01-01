package dock

import (
	"time"

	log "github.com/sirupsen/logrus"
	ui "github.com/gizak/termui/v3"
	cpb "github.com/SailGame/GoDock/pb/core"
)

type page string;
const (
	Lobby page = "lobby"
	Room  page = "room"
	Game  page = "game"
)

type Dock struct {
	UIEventC <-chan ui.Event
	CoreMsgEventC chan *cpb.BroadcastMsg
	TimeTickC chan time.Time
	Grid *ui.Grid
}

func NewDock(pollUIEvent <-chan ui.Event) *Dock {
	d := &Dock{
		UIEventC: pollUIEvent,
		CoreMsgEventC: make(chan *cpb.BroadcastMsg),
		TimeTickC: make(chan time.Time),
		Grid: ui.NewGrid(),
	}
	termWidth, termHeight := ui.TerminalDimensions()
	d.Grid.SetRect(0, 0, termWidth, termHeight)
	return d
}

func (d *Dock) Loop(){
	select{
		case e := <-d.UIEventC:
			log.Debugf("Recv UI event: %v", e)
			switch e.ID { // event string/identifier
			case "q", "<C-c>": // press 'q' or 'C-c' to quit
				return
			case "<MouseLeft>":
				// payload := e.Payload.(ui.Mouse)
				// x, y := payload.X, payload.Y
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				d.Grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				d.TimeTick()
			}
			switch e.Type {
			case ui.KeyboardEvent: // handle all key presses
				// eventID = e.ID // keypress string
			}
		case <-d.CoreMsgEventC:
			log.Debugf("Recv Core msg event: %v", d)
		// use Go's built-in tickers for updating and drawing data
		case <-d.TimeTickC:
			d.TimeTick()
	}
}

func (d *Dock) TimeTick(){
	ui.Render(d.Grid)
}

func (d *Dock) Navigate(p page){
	if p == Lobby {
		d.Grid.Set(
			ui.NewRow(1.0/2,
				ui.NewCol(1.0/2),
				ui.NewCol(1.0/2),
			),
		)
	}else if p == Room {

	}else if p == Game {
		
	}
}