package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var homeTemplate = template.Must(template.ParseFiles("home.html"))

type SocketServer struct {
	port  int
	route string
}

func newSocketServer(port int, route string) *SocketServer {
	return &SocketServer{
		port:  port,
		route: route,
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func (s *SocketServer) start() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc(s.route, func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
