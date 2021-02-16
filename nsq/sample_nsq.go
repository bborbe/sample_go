package main

import (
	"log"
	"github.com/bitly/go-nsq"
	"sync"
	"context"
	"time"
)

const topic = "test"
const channel = "test"
const address = "127.0.0.1:4150"
const timeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go produceMessages(ctx, wg)

	wg.Add(1)
	go consumerMessages(ctx, wg)

	wg.Wait()
}

func produceMessages(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(address, config)
	if err != nil {
		log.Panic("%v", err)
	}
	defer producer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := producer.Publish(topic, []byte("hello world")); err != nil {
				log.Panic("Could not connect")
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func consumerMessages(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Panic("%v", err)
	}
	defer consumer.Stop()
	log.Print("add handler")
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", message.Body)
		return nil
	}))
	if err := consumer.ConnectToNSQD(address); err != nil {
		log.Panic("Could not connect")
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
