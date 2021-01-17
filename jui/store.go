package jui

import (
	cpb "github.com/SailGame/GoDock/pb/core"
)

type Store interface {
	SetRouter(Router)
	SetGameCoreClient(cpb.GameCoreClient)
	GetRouter() Router
	GetGameCoreClient() cpb.GameCoreClient
	SetToken(string)
	GetToken() string

	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type DefaultStore struct {
	mRouter Router
	mGameCoreClient cpb.GameCoreClient
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

func (ds *DefaultStore) SetGameCoreClient(gcc cpb.GameCoreClient){
	ds.mGameCoreClient = gcc
}

func (ds *DefaultStore) GetRouter() Router{
	return ds.mRouter
}

func (ds *DefaultStore) GetGameCoreClient() cpb.GameCoreClient{
	return ds.mGameCoreClient
}

func (ds *DefaultStore) SetToken(t string){
	ds.mToken = t
}

func (ds *DefaultStore) GetToken() string{
	return ds.mToken
}

func (ds *DefaultStore) Set(key string, value interface{}){
	ds.mData[key] = value
}

func (ds *DefaultStore) Get(key string) (interface{}, bool){
	v, ok := ds.mData[key]
	return v, ok
}

