package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main(){
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err == nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil{
		log.Println("Error creating jetstream context")
	}
	
	// Simple Async Ephemeral Consumer
	_, err = js.Subscribe("jet", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})

	if err != nil {
		log.Println("Subscription failed!")
	}

	// Keep the connection alive
	runtime.Goexit()
}