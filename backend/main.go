package main

import (
	"fmt"
	"net/http"

	"github.com/shoebilyas123/gochat/pkg/websocket"
)

func serveWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host);

	ws, err := websocket.Upgrader.Upgrade(w,r,nil);

	if err != nil {
		fmt.Println(err);
		return;
	}

	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Simple server")
	})
	http.HandleFunc("/ws", serveWS)
}

func main() {
	fmt.Println("Hello websockets")
	setupRoutes()
	http.ListenAndServe(":8080",nil)
}