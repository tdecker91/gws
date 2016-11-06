package main

import (
	"flag"
	"log"
)

var port = flag.Int("port", 8080, "port to listen on")
var route = flag.String("route", "/ws", "route to listen for socket connections")

func main() {
	flag.Parse()
	log.Println("Starting websocket server...")
	log.Printf("Listening on port %d for '%s'", *port, *route)

	server := newSocketServer(*port, *route)
	server.start()
}
