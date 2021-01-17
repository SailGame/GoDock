package conn

import (
	"context"
	"errors"

	cpb "github.com/SailGame/GoDock/pb/core"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GameCoreConn struct {
	mCoreClient cpb.GameCoreClient
	mLisClient cpb.GameCore_ListenClient
	mCancel    func()
	mMsgC	chan *cpb.BroadcastMsg
	mToken string
}

func NewGameCoreConn(grpcConn *grpc.ClientConn) *GameCoreConn {
	return &GameCoreConn{
		mCoreClient: cpb.NewGameCoreClient(grpcConn),
		mMsgC: make(chan *cpb.BroadcastMsg),
	}
}

func (uc *GameCoreConn) Login(userName string) error {
	loginRet, err := uc.mCoreClient.Login(context.TODO(), &cpb.LoginArgs{
		UserName: userName,
	})
	if err != nil{
		return err
	}
	if loginRet.Err != cpb.ErrorNumber_OK{
		return errors.New(loginRet.GetErr().String())
	}
	uc.mToken = loginRet.Token
	return nil
}

func (uc *GameCoreConn) CloseListenStream() error {
	if(uc.mCancel == nil){
		return errors.New("No live listen stream")
	}
	uc.mCancel()
	uc.mCancel = nil
	return nil
}

func (uc *GameCoreConn) ListenToCore() error {
	ctx, cancel := context.WithCancel(context.Background())
	var err error
	uc.mLisClient, err = uc.mCoreClient.Listen(ctx, &cpb.ListenArgs{
		Token: uc.mToken,
	})
	if err != nil {
		return err
	}
	uc.mCancel = cancel
	return err
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

func (uc *GameCoreConn) GetBroadcastMsgCh() <-chan *cpb.BroadcastMsg {
	return uc.mMsgC
}

func (uc *GameCoreConn) GetToken() string {
	return uc.mToken
}

func (uc *GameCoreConn) GetGameCoreClient() cpb.GameCoreClient {
	return uc.mCoreClient
}