package main

import (
	"log"
)

type Hub struct {
	clients    map[string]*client
	fromClient chan Message
	toClient   chan Message
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func newHub() *Hub {
	return &Hub{
		fromClient: make(chan Message),
		toClient:   make(chan Message),
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[string]*client),
	}
}

func (h *Hub) closeConnection(clientId string) {
	close(h.clients[clientId].send)
	delete(h.clients, clientId)
	log.Printf("Closing connection to client %v\n", clientId)
}

func (h *Hub) run(outbound chan Message) {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
			outbound <- *(newMessage(client.id, ClientConnected, []byte{}))
		case client := <-h.unregister:
			if _, ok := h.clients[client.id]; ok {
				h.closeConnection(client.id)
				outbound <- *(newMessage(client.id, ClientDisconnected, []byte{}))
			}
		case message := <-h.fromClient:
			outbound <- message

		case message := <-h.toClient:
			select {
			case h.clients[message.ClientId].send <- message.Data:
			default:
				h.closeConnection(message.ClientId)
			}

		case broadcast := <-h.broadcast:
			for clientId := range h.clients {
				select {
				case h.clients[clientId].send <- broadcast:
				default:
					h.closeConnection(clientId)
				}
			}
		}
	}
}
