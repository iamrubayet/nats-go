package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	// connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// simple async subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received a message: %s", string(m.Data))
	})

	// simple publisher
	nc.Publish("foo", []byte("Hello World"))

	// flush connection to server, returns when all messages have been processed
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Published message on subject 'foo'")

	// keep the connection alive
	select {}

}
