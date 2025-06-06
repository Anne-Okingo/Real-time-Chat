package main

import(
	"golang.org/x/net/websocket"
	"fmt"
)

//ChatServer keeps track of all clients connected to the websocket connection
type ChatServer struct{
clients map[*websocket.Conn]bool
}

//this is a new chat server constructor, it also initializes the websocket
func NewChatServer() *ChatServer {
	return &ChatServer {
		clients: make(map[*websocket.Conn]bool),
	}
}


//handleConnection function that manages a new websocket connection eg broadcast msgs
func (cs * ChatServer) hnadleConnection(ws *websocket.Conn) {
	//Add new client to the list
	clientAddr := ws.RemoteAddr().String()
	fmt.Printf("Anew client connected:  %s/n", clientAddr)
	cs.clients[ws] = true
}