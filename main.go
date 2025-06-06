package main

import(
	"golang.org/x/net/websocket"
)

//ChatServer keeps track of all clients connected to the websocket connection
type ChatServer struct{
clients map[*websocket.Conn]bool
}

//this is a new chat server constructor, it also initializes the websocket
func NewChatServer() * ChatServer {
	return &ChatServer {
		clients: make(map[*websocket.Conn]bool),
	}
}