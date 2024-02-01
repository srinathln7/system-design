package main

import (
	chat "chat/internal/chatroom"
	"chat/internal/feed"
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	cr := chat.NewChatRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", cr)

	// get the room going
	go cr.Run()

	// get the feed going
	var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	http.HandleFunc("/feed", func(w http.ResponseWriter, req *http.Request) {
		socket, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Fatal("error starting feed:", err)
			return
		}

		feed := feed.NewFeed(socket)
		feed.GenerateFeedToUser()
	})

	// start the web server
	log.Println("starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("listen and serve:", err)
	}

}
