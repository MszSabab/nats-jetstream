package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Println("Error creating stream context!")
	}
	
	// Simple Jetstream Publisher
	_, err = js.Publish("jet", []byte("Hello World"))
	if err != nil {
		log.Println("Message not published")
	}	

	defer nc.Close()
}