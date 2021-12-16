package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("hello there")
	http.HandleFunc("/ws", websocketHandler)
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe(": 8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	for {
		if err := conn.WriteMessage(1, []byte("helllo there")); err != nil {
			log.Println(err)
			return
		}
		mtype, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}
		fmt.Println(mtype, string(p))
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello we aare fukkkd up")
}

// u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
