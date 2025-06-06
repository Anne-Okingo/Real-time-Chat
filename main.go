package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/websocket"
)

// ChatServer keeps track of all clients connected to the websocket connection
type ChatServer struct {
	clients map[*websocket.Conn]bool
}

// this is a new chat server constructor, it also initializes the websocket
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[*websocket.Conn]bool),
	}
}

// handleConnection function that manages a new websocket connection eg broadcast msgs
func (cs *ChatServer) handleConnection(ws *websocket.Conn) {
	// Add new client to the list
	clientAddr := ws.RemoteAddr().String()
	fmt.Printf("Anew client connected:  %s/n", clientAddr)
	cs.clients[ws] = true

	// Ensure clean up after client is disconnected
	defer func() {
		ws.Close()
		delete(cs.clients, ws)
		fmt.Printf("Client from: %s disconnected \n", clientAddr)
	}()

	// start listening for maessages from clients

	for {
		messages := ""

		err := websocket.Message.Receive(ws, &messages)
		if err != nil {
			// client disconnected or error occurred
			break
		}

		fmt.Printf("Received message from %s\n", clientAddr, messages)

		// broadcast the message to all clients
		cs.broadcast(messages)
	}
}

// broacdcast send the websocket message to all connected clients
func (cs *ChatServer) broadcast(message string) {
	for client := range cs.clients {
		err := websocket.Message.Send(client, message)
		if err != nil {
			// if sending fails, assume client disconnected
			fmt.Println("Error sending to client:", err)
			client.Close()
			delete(cs.clients, client)
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		ws.Config().Origin, _ = url.Parse("http://localhost")
		fmt.Println("Client connected")

		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				break
			}
			fmt.Println("Received:", msg)
			websocket.Message.Send(ws, "Server: "+msg)
		}
	}))

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
