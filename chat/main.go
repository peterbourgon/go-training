package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

// Based on The Go Programming Language chapter 8 "chat" example.
// https://github.com/adonovan/gopl.io/blob/master/ch8/chat/chat.go

func main() {
	var (
		server = flag.Bool("server", false, "server mode")
		client = flag.Bool("client", false, "client mode")
		addr   = flag.String("addr", "127.0.0.1:5000", "chat server address")
	)
	flag.Parse()

	switch {
	case *server && *client:
		log.Fatal("can't be both server and client")
	case *server && !*client:
		runServer(*addr)
	case !*server && *client:
		runClient(*addr)
	case !*server && !*client:
		log.Fatal("must be either server or client")
	}
}

// To the server, a client is a pipe that needs to receive messages.
type client chan<- string

func runServer(addr string) {
	// Bind listener
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server: listening on %s", ln.Addr().String())

	// Client lifecycle management
	var (
		hello     = make(chan client) // client connects
		goodbye   = make(chan client) // client disconnects
		broadcast = make(chan string) // messages to all clients
	)

	// The goroutine to send broadcasts.
	go func() {
		clients := map[client]struct{}{}
		for {
			select {
			case client := <-hello:
				log.Printf("hello: now %d client(s)", len(clients))
				clients[client] = struct{}{}

			case client := <-goodbye:
				log.Printf("goodbye: now %d client(s)", len(clients))
				delete(clients, client)
				close(client)

			case message := <-broadcast:
				log.Printf("broadcast: to %d client(s)", len(clients))
				for client := range clients {
					client <- message
				}
			}
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleClient(conn, hello, goodbye, broadcast)
	}
}

func handleClient(conn net.Conn, hello, goodbye chan<- client, toBroadcast chan<- string) {
	// Lifecycle management.
	toClient := make(chan string)
	hello <- toClient
	defer func() { goodbye <- toClient }()

	// Announce the presence of a new chatter.
	handle := conn.RemoteAddr().String()
	toBroadcast <- fmt.Sprintf("%s has joined the chat", handle)
	defer func() { toBroadcast <- fmt.Sprintf("%s has left the chat", handle) }()

	// Messages from the chat server to the conn.
	// We range over the toClient channel and exit when it's closed.
	// (It's closed after it's processed by the goodbye handler.)
	go func() {
		for message := range toClient {
			fmt.Fprintf(conn, message+"\n")
		}
	}()

	// Messages from the conn to the chat server.
	// When the client disconnects, the scanner stops scanning.
	// We close the conn, and signal our goodbye via the defer.
	s := bufio.NewScanner(conn)
	for s.Scan() {
		toBroadcast <- "<" + handle + "> " + s.Text()
	}
}

func runClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Messages from the conn to stdout.
	// The scanner will exit when the conn is closed.
	// We close the conn via the defer.
	go func() {
		s := bufio.NewScanner(conn)
		for s.Scan() {
			fmt.Fprintf(os.Stdout, s.Text()+"\n")
		}
	}()

	// Messages from stdin to the conn.
	// The scanner will exit when stdin is closed.
	// We close stdin via ctrl-D.
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Fprintf(conn, s.Text()+"\n")
	}
}
