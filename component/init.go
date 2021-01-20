package component

import "github.com/SailGame/GoDock/jui"

func Init(store jui.Store){
	router := store.GetRouter()
	router.RegisterComponent(LOBBY, NewLobby(NewDefaultComponent(store)))
	router.RegisterComponent(ROOM, NewRoom(NewDefaultComponent(store)))
}