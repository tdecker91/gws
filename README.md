# Simple Golang Websocket Server
This project was created to take the hassle out of setting up a websocket server in golang. It's simple enough to be used in any project.

## Usage
install the gws package `go get github.com/tdecker91/gws`

Create a new socket server and provide a port to listen on and the route to listen 
```
  server := gws.NewSocketServer(8080, "/ws")
```

Starting the server requires a channel for sending messages to.
```
  messages := make(chan gws.Message)
  go server.Start(messages)
```

Recieve messages from the channel
```
  msg := <-messages
```

## Echo Example
```
package main

import (
	"flag"
	"log"

	"github.com/tdecker91/gws"
)

var port = flag.Int("port", 8080, "port to listen on")
var route = flag.String("route", "/ws", "route to listen for socket connections")

func main() {
	flag.Parse()
	log.Println("Starting websocket server...")
	log.Printf("Listening on port %d for '%s'", *port, *route)

	messages := make(chan gws.Message)

	server := gws.NewSocketServer(*port, *route)
	go server.Start(messages)

	for {
		m := <-messages

		switch t := m.Type; t {
		case gws.ClientConnected:
			log.Printf("Client connected %v\n", m.ClientId)
		case gws.ClientDisconnected:
			log.Printf("Client disconnected %v", m.ClientId)
		case gws.ClientMessage:
			log.Printf("Client message received %v", m.Data)
			server.SendMessage(m)
		}
	}
}

```
