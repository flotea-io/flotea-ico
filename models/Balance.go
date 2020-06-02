/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

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
