package main

import (
	"log"
	"net"
)

const (
	PORT        string = ":8080"
	BUFFER_SIZE int    = 1024
)

func main() {

	// Create a tcp listener
	ln, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error creating tcp server: %s", err.Error())
	}

	log.Printf("server started on port: %s \n", PORT)
	defer ln.Close()

	for {

		// Listen for incoming connections
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("error accepting incoming client connection: %s", err.Error())
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	// Make a buffer to temp. store incoming msges from clients

	buf := make([]byte, BUFFER_SIZE)

	n, err := conn.Read(buf)
	if err != nil {
		log.Println("failed to read from connection")
		return
	}

	_, err = conn.Write(buf[:n])
	if err != nil {
		log.Println("failed to write to connection")
		return
	}

	log.Printf("Received data from %v: %s \n", conn.RemoteAddr(), string(buf[:n]))
}
