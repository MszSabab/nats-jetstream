package main

import (
	"log"
	"runtime"

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
	if err != nil{
		log.Println("Error creating jetstream context")
	}
	
	// Simple Async Ephemeral Consumer
	js.Subscribe("jet", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})
	
	// // Simple Pull Consumer
	// sub, err := js.PullSubscribe("jet", "consumer1")
	// if err != nil {
	// 	panic("Could not subscribe")
	// }

	// msgs, err := sub.Fetch(10)

	// if err == nil {
	// 	log.Println(msgs)
	// }

	// Simple Sync Durable Consumer (optional SubOpts at the end)
	// timeout := 10*time.Second
	
	// // Simple Async Ephemeral Consumer
	// js.Subscribe("jet", func(msg *nats.Msg) {
	// 	log.Println(string(msg.Data))
	// })

	// sub, err := js.SubscribeSync("jet", nats.Durable("MONITOR"), nats.MaxDeliver(3))
	// if err != nil {
	// 	log.Println("Could not subscribe")
	// }

	// msg, err := sub.NextMsg(timeout)
	// if err == nil {
	// 	log.Println(string(msg.Data))
	// }


	// // Simple Async Subscriber
	// nc.Subscribe("foo", func(msg *nats.Msg) {
	// 	log.Println(string(msg.Data))
	// })

	// Keep the connection alive
	runtime.Goexit()
}