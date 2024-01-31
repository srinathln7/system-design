package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type chatroom struct {
	agents map[*agent]bool

	join      chan *agent
	leave     chan *agent
	broadcast chan []byte
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (cr *chatroom) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	agent := &agent{
		socket:   socket,
		recvCh:   make(chan []byte, messageBufferSize),
		chatroom: cr,
	}
	cr.join <- agent
	defer func() { cr.leave <- agent }()

	go agent.sendMsgToChatRoom()
	agent.deliverMsgToUser()
}

func NewChatRoom() *chatroom {
	return &chatroom{
		agents:    make(map[*agent]bool),
		join:      make(chan *agent),
		leave:     make(chan *agent),
		broadcast: make(chan []byte),
	}
}

func (cr *chatroom) Run() {
	for {
		select {
		case agent := <-cr.join:
			cr.agents[agent] = true
		case agent := <-cr.leave:
			delete(cr.agents, agent)
		case payload := <-cr.broadcast:
			// broadcast the msg to the rest of the agents in the chat room
			for agent := range cr.agents {
				agent.recvCh <- payload
			}
		}
	}
}
