package main

import(
	"golang.org/x/net/websocket"
)

//ChatServer keeps track of all clients connected to the websocket connection
type ChatServer struct{
clients map[*websocket.Conn]bool

}