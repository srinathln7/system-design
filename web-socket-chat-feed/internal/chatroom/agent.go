package chat

import (
	"github.com/gorilla/websocket"
)

// agent: Represents an unique entity/service maintaining a web-socket connection to the end-user (a.k.a) client
// Rem: each client(user) maintains an individual web socket connection with its client
type agent struct {
	chatroom *chatroom
	socket   *websocket.Conn
	recvCh   chan []byte
}

// sendMsgFromUser: agent continously listens for msg on the websocket from the end user
// and forwards it to the chat room
func (agent *agent) sendMsgToChatRoom() {
	defer agent.socket.Close()

	// read the msg that the client (user) sends to the ws
	// and forward it to the chat room
	for {
		// payload -> msg in bytes
		_, payload, err := agent.socket.ReadMessage()
		if err != nil {
			continue
		}
		agent.chatroom.broadcast <- payload
	}
}

// deliverMsgToUser: agent receives the msg from other agents in the chat room
// and delivers to the user via websocket
func (agent *agent) deliverMsgToUser() {
	defer agent.socket.Close()

	// payload -> msg in bytes
	for payload := range agent.recvCh {
		err := agent.socket.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			continue
		}
	}
}
