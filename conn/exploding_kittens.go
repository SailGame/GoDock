package conn

import (
)

//go:generate mockgen -destination=mocks/exploding_kittens.go -package=mocks . ExplodingKittensClient
type ExplodingKittensClient interface{
	Swap() (error)
}

func (gcc *GameCoreConn) Swap() error {
	return nil
}