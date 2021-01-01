package conn

import (
	"context"
	"errors"

	cpb "github.com/SailGame/GoDock/pb/core"
	log "github.com/sirupsen/logrus"
)

type GameCoreConn struct {
	mCoreClient cpb.GameCoreClient
	mLisClient cpb.GameCore_ListenClient
	mCancel    func()
	mMsgC		chan *cpb.BroadcastMsg
}

func (uc *GameCoreConn) CloseListenStream(token string) (err error) {
	if(uc.mCancel == nil){
		return errors.New("No live listen stream")
	}
	uc.mCancel()
	uc.mCancel = nil
	return nil
}

func (uc *GameCoreConn) ListenToCore(token string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	uc.mLisClient, err = uc.mCoreClient.Listen(ctx, &cpb.ListenArgs{
		Token: token,
	})
	if err != nil {
		return
	}
	uc.mCancel = cancel
	return
}

func (uc *GameCoreConn) LoopListenStream(onStop func()){
	for{
		msg, err := uc.mLisClient.Recv()
		if err != nil {
			log.Warn(err)
			onStop()
			return
		}
		log.Debugf("GameCoreConn Recv Msg %v", msg)
		uc.mMsgC <- msg
	}
}