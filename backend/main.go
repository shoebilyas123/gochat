package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {return true},
}

func reader(conn *websocket.Conn) {
	for {
		mType, msg, err := conn.ReadMessage();

		if err != nil {
			log.Fatal(err)
			return;
		}

		fmt.Println(string(msg))

		if err := conn.WriteMessage(mType, msg); err != nil {
			log.Fatal(err);
			return;
		}
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host);

	ws, err := upgrader.Upgrade(w,r,nil);

	if err != nil {
		fmt.Println(err);
		return;
	}

	reader(ws)
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