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
	js, err := nc.JetStream()
	if err != nil {
		log.Println("Error creating stream context!")
	}

	// // Create a Stream
	// _, err = js.AddStream(&nats.StreamConfig{
	// 	Name:     "stream1",
	// 	Subjects: []string{"jet"},
	// })

	// if err != nil {
	// 	log.Println("Error adding stream!")
	// }

	// // Create a Consumer
	// _, err = js.AddConsumer("stream1", &nats.ConsumerConfig{
	// 	Durable: "consumer1",
	// })

	// if err != nil {
	// 	log.Println("Error adding consumer!")
	// }

	// js := jetstreamManager(nc)
	
	// Simple Jetstream Publisher
	_, err = js.Publish("jet", []byte("Hello World"))
	if err != nil {
		log.Println("Message not published")
	}	
	
	// // Create JetStream Context
	// js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	// if err != nil {
	// 	log.Println("Could not create jetstream context")
	// }

	// // Simple Async Stream Publisher
	// for i := 0; i < 500; i++ {
	// 	js.PublishAsync("jets", []byte("hello"))
	// }
	// select {
	// 	case <-js.PublishAsyncComplete():
	// 	case <-time.After(5 * time.Second):
	// 		log.Println("Did not resolve in time")
	// }

	defer nc.Close()
}

// func jetstreamManager(nc *nats.Conn) nats.JetStreamContext{
// 	// Create JetStream Context
// 	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

// 	// Create a Stream
// 	_, err := js.AddStream(&nats.StreamConfig{
// 		Name:     "stream1",
// 		Subjects: []string{"jet"},
// 	})

// 	if err != nil {
// 		log.Println("Error adding stream!")
// 	}

// 	// Create a Consumer
// 	_, err = js.AddConsumer("stream1", &nats.ConsumerConfig{
// 		Durable: "consumer1",
// 	})

// 	if err != nil {
// 		log.Println("Error adding consumer!")
// 	}

// 	return js
// }