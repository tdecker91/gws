package main

import (
	"log"
	"net/http"
	"strconv"
)

type SocketServer struct {
	port  int
	route string
	hub   *Hub
}

func newSocketServer(port int, route string) *SocketServer {
	return &SocketServer{
		port:  port,
		route: route,
		hub:   nil,
	}
}

func messageHandler(msg interface{}, hub *Hub, fn func(interface{})) {
	if hub == nil {
		log.Fatalln("Tried sending a message before starting the socket server")
		return
	}

	fn(msg)
}

func (s *SocketServer) SendMessage(msg Message) {
	messageHandler(msg, s.hub, func(m interface{}) {
		s.hub.toClient <- m.(Message)
	})
}

func (s *SocketServer) BroadcastMessage(msg []byte) {
	messageHandler(msg, s.hub, func(m interface{}) {
		s.hub.broadcast <- m.([]byte)
	})
}

func (s *SocketServer) Start(outbound chan Message) {
	s.hub = newHub()
	go s.hub.run(outbound)

	http.HandleFunc(s.route, func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	})

	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(s.port), nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
