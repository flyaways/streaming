package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/flyaways/streaming"
)

func Processor(msg *sarama.ConsumerMessage, outTopic string) (*sarama.ProducerMessage, error) {
	return &sarama.ProducerMessage{
		Topic: outTopic,
		Key:   sarama.ByteEncoder(msg.Key),
		Value: sarama.ByteEncoder(msg.Value),
	}, nil
}

func main() {
	if err := streaming.NewStreaming(
		[]string{"127.0.0.1:9092"},
		[]string{"input-topic", "input-topic-2"},
		"streaming-group",
		"output-topic",
		cluster.NewConfig(),
		Processor); err != nil {
		log.Println(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}
