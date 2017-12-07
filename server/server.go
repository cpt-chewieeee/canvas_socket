package main

import (
	"fmt"
	"net/http"
	// "github.com/julienschmidt/httprouter"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	// "time"
)
type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {
	fmt.Println("_____API_SERVICES_STARTED_____")

	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("connected")

		c.Emit("/message", Message{10, "main", "using emit"})
	})
	server.On("/join", func(c *gosocketio.Channel, channel Channel) string {
		log.Println("Client joined to ", channel.Channel)
		return "joined to " + channel.Channel
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("starting server...")
	log.Panic(http.ListenAndServe(":3131", serveMux))
}