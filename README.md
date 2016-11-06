# Simple Golang Websocket Server
This project was created to take the hassle out of setting up a websocket server in golang. It's simple enough to be used in any project.

## Usage
install the gws package `go get github.com/tdecker91/gws`

Create a new socket server and provide a port to listen on and the route to listen 
```go
  server := gws.NewSocketServer(8080, "/ws")
```

Starting the server requires a channel for sending messages to.
```go
  messages := make(chan gws.Message)
  go server.Start(messages)
```

Recieve messages from the channel. Data is received as a byte array, so decode it however you intend it. (json etc...)
```go
  msg := <-messages
  
  if msg.Type == gws.ClientMessage {
  	fmt.Printf("Received Message from client %s: %v\n", msg.ClientId, msg.Data)
  }
```

Send messages to a single client, or broadcase messages to all clients
```go
  server.BroadcastMessage([]byte{`This will be sent to all clients`})
  server.SendMessage(gws.NewMessage(clientId, []byte{`This will only be sent to a single client`}))
```

And that's it. If there are any suggestions or issues feel free to contribute, or open an issue on github.

## Echo Example
```go
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
			log.Printf("Client connected %s\n", m.ClientId)
		case gws.ClientDisconnected:
			log.Printf("Client disconnected %s", m.ClientId)
		case gws.ClientMessage:
			log.Printf("Client message received\n\tClient: t%s\n\tMessage: %v", m.ClientId, m.Data)
			server.SendMessage(m)
		}
	}
}

```
