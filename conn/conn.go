package conn

import (
	cpb "github.com/SailGame/GoDock/pb/core"
)

//go:generate mockgen -destination=mocks/conn.go -package=mocks . GameCoreClient
type GameCoreClient interface{
	TexasClient
	ExplodingKittensClient

	// client
	Login(in *cpb.LoginArgs) (*cpb.LoginRet)
	QueryAccount(in *cpb.QueryAccountArgs) (*cpb.QueryAccountRet)
	// room
	CreateRoom(in *cpb.CreateRoomArgs) (*cpb.CreateRoomRet)
	ControlRoom(in *cpb.ControlRoomArgs) (*cpb.ControlRoomRet)
	ListRoom(in *cpb.ListRoomArgs) (*cpb.ListRoomRet)
	JoinRoom(in *cpb.JoinRoomArgs) (*cpb.JoinRoomRet)
	ExitRoom(in *cpb.ExitRoomArgs) (*cpb.ExitRoomRet)
	QueryRoom(in *cpb.QueryRoomArgs) (*cpb.QueryRoomRet)
	// OperationInRoom
	Ready(in *cpb.OperationInRoomArgs_Ready) (*cpb.OperationInRoomRet)
}