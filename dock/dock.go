package dock

import (
	"time"

	log "github.com/sirupsen/logrus"
	ui "github.com/gizak/termui/v3"
	cpb "github.com/SailGame/GoDock/pb/core"
)

type Dock struct {
	mGameCoreClient cpb.GameCoreClient
	mUIEventC <-chan ui.Event
	mCoreMsgEventC <-chan *cpb.BroadcastMsg
	mTimeTickC <-chan time.Time
	mRouter Router
}

func NewDock(gameCoreClient cpb.GameCoreClient, pollUIEvent <-chan ui.Event, coreMsgEventC <-chan *cpb.BroadcastMsg, timeTickC <-chan time.Time) *Dock {
	d := &Dock{
		mGameCoreClient: gameCoreClient,
		mUIEventC: pollUIEvent,
		mCoreMsgEventC: coreMsgEventC,
		mTimeTickC: timeTickC,
		mRouter: NewDefaultRouter(map[string]Component{
			"/": NewLobbyComponent(gameCoreClient),
			// "/room": NewRoomComponent(),
		}),
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
				d.mRouter.GetCurrentComponent().GetGrid().SetRect(0, 0, payload.Width, payload.Height)
			case "<F1>":
				d.mRouter.NavigateBack()
			default:
				d.mRouter.GetCurrentComponent().HandleUIEvent(e)
			}
		case e := <-d.mCoreMsgEventC:
			log.Debugf("Recv Core msg event: %v", d)
			d.mRouter.GetCurrentComponent().HandleServerEvent(e)
		case <-d.mTimeTickC:
			log.Debugf("TimeTick")
			d.mRouter.GetCurrentComponent().TimeTick()
			d.TimeTick()
		}
	}
}

func (d *Dock) TimeTick(){
	// termWidth, termHeight := ui.TerminalDimensions()
	// d.Grid.SetRect(0, 0, termWidth, termHeight)
	ui.Render(d.mRouter.GetCurrentComponent().GetGrid())
}