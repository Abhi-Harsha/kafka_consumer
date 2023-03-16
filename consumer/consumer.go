package consumer

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Read(c *kafka.Consumer) {
	c.SubscribeTopics([]string{"first_topic"}, nil)
	fmt.Println("connected....")
	// A signal handler or similar could be used to set this to false to break the loop.
	flag := true
	for flag {
		msg, err := c.ReadMessage(time.Millisecond * 10)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
