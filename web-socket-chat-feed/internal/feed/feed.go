package feed

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type feed struct {
	socket *websocket.Conn
}

func NewFeed(socket *websocket.Conn) *feed {
	return &feed{
		socket: socket,
	}
}

func (feed *feed) GenerateFeedToUser() {
	defer feed.socket.Close()

	var payload []byte
	for {
		payload = []byte(fmt.Sprintf("live feed received at %d\n", time.Now().UnixNano()))
		err := feed.socket.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond)
	}
}
