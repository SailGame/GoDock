package dock

import (
	"fmt"
	"time"

	"github.com/SailGame/GoDock/jui"
	cpb "github.com/SailGame/GoDock/pb/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	log "github.com/sirupsen/logrus"
)

type Dock struct {
	mStore jui.Store
	mUIEventC <-chan ui.Event
	mCoreMsgEventC <-chan *cpb.BroadcastMsg
	mTimeTickC <-chan time.Time
	mGrid *ui.Grid
}

func NewDock(store jui.Store, pollUIEvent <-chan ui.Event, coreMsgEventC <-chan *cpb.BroadcastMsg, timeTickC <-chan time.Time) *Dock {
	d := &Dock{
		mStore: store,
		mUIEventC: pollUIEvent,
		mCoreMsgEventC: coreMsgEventC,
		mTimeTickC: timeTickC,
		mGrid: ui.NewGrid(),
	}
	termWidth, termHeight := ui.TerminalDimensions()
	d.mGrid.SetRect(0, 0, termWidth, termHeight)
	return d
}

func (d *Dock) Loop(){
	for {
		select{
		case e := <-d.mUIEventC:
			log.Debugf("Dock Recv UI event: %v", e)
			switch e.ID { // event string/identifier
			case "<C-c>": // press 'C-c' to quit
				log.Info("Dock Received C-c. Closing Dock")
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				d.mGrid.SetRect(0, 0, payload.Width, payload.Height)
			case "<F1>":
				d.mStore.GetRouter().NavigateBack()
			default:
				d.GetCurrentComponent().HandleUIEvent(e)
			}
		case e := <-d.mCoreMsgEventC:
			log.Debugf("Recv Core msg event: %v", d)
			d.GetCurrentComponent().HandleServerEvent(e)
		case <-d.mTimeTickC:
			d.TimeTick()
		}
	}
}

func (d *Dock) TimeTick(){
	// termWidth, termHeight := ui.TerminalDimensions()
	// d.Grid.SetRect(0, 0, termWidth, termHeight)
	d.GetCurrentComponent().TimeTick()
	breadcrumb := widgets.NewParagraph()
	breadcrumb.Text = fmt.Sprintf("CurrentPage: %s", d.mStore.GetRouter().GetCurrentPath()) 
	d.mGrid.Set(
		ui.NewRow(0.2/2, ui.NewCol(1, breadcrumb)),
		ui.NewRow(1.0/2, ui.NewCol(1, d.GetCurrentComponent().GetGrid())),
	)
	ui.Render(d.mGrid)
}

func (d *Dock) GetCurrentComponent() jui.Component{
	return d.mStore.GetRouter().GetCurrentComponent()
}