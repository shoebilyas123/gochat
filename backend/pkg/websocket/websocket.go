package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {return true},
}

func Reader(conn *websocket.Conn) {
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
