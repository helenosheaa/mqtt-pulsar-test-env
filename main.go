package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func main() {
	// Instantiate a Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})

	if err != nil {
		fmt.Println("Could not create client:", err)
	}

	defer client.Close()

	// Use the client to instantiate a producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "temperature",
	})
	if err != nil {
		fmt.Println("Could not create producer:", err)
	}

	ctx := context.Background()
	// Send 10 messages synchronously and 10 messages asynchronously
	for i := 0; i < 10; i++ {
		// Create a message
		msg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("message-%d", i)),
		}
		// Attempt to send the message
		if err := producer.Send(ctx, msg); err != nil {
			fmt.Println("Could not send message:", err)
		}
		fmt.Printf("the %s successfully published", string(msg.Payload))
		// Create a different message to send asynchronously
		asyncMsg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("async-message-%d", i)),
		}
		// Attempt to send the message asynchronously and handle the response
		producer.SendAsync(ctx, asyncMsg, func(msg pulsar.ProducerMessage, err error) {
			if err != nil {
				fmt.Println("Could not send message async:", err)
			}
			fmt.Printf("the %s successfully published", string(msg.Payload))
		})
	}
	defer producer.Close()

	// Consumer channel
	msgChannel := make(chan pulsar.ConsumerMessage)
	consumerOpts := pulsar.ConsumerOptions{
		Topic:            "temperature",
		SubscriptionName: "my-subscription-1",
		Type:             pulsar.Exclusive,
		MessageChannel:   msgChannel,
	}
	consumer, err := client.Subscribe(consumerOpts)
	if err != nil {
		fmt.Println("Could not create subscription:", err)
	}
	defer consumer.Close()
	for cm := range msgChannel {
		msg := cm.Message
		fmt.Printf("Message ID: %s", msg.ID())
		fmt.Printf("Message value: %s", string(msg.Payload()))
		consumer.Ack(msg)
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}
