package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/flyaways/streaming"
)

func Processor(msg *sarama.ConsumerMessage, outTopic []string) ([]*sarama.ProducerMessage, error) {
	msgs := []*sarama.ProducerMessage{}
	if msg.Topic == "input-topic-2" {
		msgs = append(msgs, &sarama.ProducerMessage{
			Topic: outTopic[0],
			Key:   sarama.ByteEncoder(msg.Key),
			Value: sarama.ByteEncoder(msg.Value),
		})
	}
	return msgs, nil
}

func main() {
	if err := streaming.NewStreaming(
		[]string{"127.0.0.1:9092"},
		[]string{"input-topic1", "input-topic-2"},
		[]string{"output-topic1", "output-topic"},
		"flyaways-streaming-group",
		cluster.NewConfig(),
		Processor); err != nil {
		log.Panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}
