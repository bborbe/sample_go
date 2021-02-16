package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
	"log"
)

func main() {
	const TOPIC = "mytopic/test"

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("connect failed: %v", token.Error())
		return
	}
	log.Printf("connected")

	var wg sync.WaitGroup
	wg.Add(1)

	if token := client.Subscribe(TOPIC, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("received msg: %s", msg.Payload())
		wg.Done()
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("subscribe failed: %v", token.Error())
	}
	log.Printf("subscribed")

	if token := client.Publish(TOPIC, 0, false, "mymessage"); token.Wait() && token.Error() != nil {
		log.Fatalf("publish failed: %v", token.Error())
		return
	}
	log.Printf("msg published")
	wg.Wait()

	log.Printf("done")
}
