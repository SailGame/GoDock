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
	mUIEventC <-chan ui.Event
	mCoreMsgEventC <-chan *cpb.BroadcastMsg
	mTimeTickC <-chan time.Time
	Grid *ui.Grid
}

func NewDock(pollUIEvent <-chan ui.Event, coreMsgEventC <-chan *cpb.BroadcastMsg, timeTickC <-chan time.Time) *Dock {
	d := &Dock{
		mUIEventC: pollUIEvent,
		mCoreMsgEventC: coreMsgEventC,
		mTimeTickC: timeTickC,
		Grid: ui.NewGrid(),
	}
	return d
}

func (d *Dock) Loop(){
	for {
		select{
		case e := <-d.mUIEventC:
			log.Debugf("Recv UI event: %v", e)
			switch e.ID { // event string/identifier
			case "<C-c>": // press 'C-c' to quit
				log.Info("Received C-c. Closing Dock")
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
		case <-d.mCoreMsgEventC:
			log.Debugf("Recv Core msg event: %v", d)
		// use Go's built-in tickers for updating and drawing data
		case <-d.mTimeTickC:
			d.TimeTick()
		}
	}
}

func (d *Dock) TimeTick(){
	// termWidth, termHeight := ui.TerminalDimensions()
	// d.Grid.SetRect(0, 0, termWidth, termHeight)
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