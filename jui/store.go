package jui

import (
	"github.com/SailGame/GoDock/conn"
)

//go:generate mockgen -destination=mocks/store.go -package=mocks . Store
type Store interface {
	SetRouter(Router)
	SetGameCoreClient(conn.GameCoreClient)
	GetRouter() Router
	GetGameCoreClient() conn.GameCoreClient

	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type DefaultStore struct {
	mRouter Router
	mGameCoreClient conn.GameCoreClient
	mToken string
	mData map[string]interface{}
}

func NewDefaultStore() *DefaultStore {
	ds := &DefaultStore{
		mRouter: nil,
		mGameCoreClient: nil,
		mData: make(map[string]interface{}),
	}
	return ds
}

func (ds *DefaultStore) SetRouter(r Router){
	ds.mRouter = r
}

func (ds *DefaultStore) SetGameCoreClient(gcc conn.GameCoreClient){
	ds.mGameCoreClient = gcc
}

func (ds *DefaultStore) GetRouter() Router{
	return ds.mRouter
}

func (ds *DefaultStore) GetGameCoreClient() conn.GameCoreClient{
	return ds.mGameCoreClient
}

func (ds *DefaultStore) Set(key string, value interface{}){
	ds.mData[key] = value
}

func (ds *DefaultStore) Get(key string) (interface{}, bool){
	v, ok := ds.mData[key]
	return v, ok
}

