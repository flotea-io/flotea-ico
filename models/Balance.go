package models

import (
	"github.com/gorilla/websocket"
)

type singleton struct {
	Balance
}

var (
	instance singleton
)

func WsData() singleton {

	once.Do(func() { // <-- atomic, does not allow repeating
		instance = singleton{Clients: make(map[*websocket.Conn]bool), Broadcast: make(chan Message)} // <-- thread safe
		//go SendBroadcasts()
	})

	return instance
}
