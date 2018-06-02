package streaming

import (
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//Streaming config
type Streaming struct {
	SeedBrokers []string
	Config      *cluster.Config
	InTopic     []string
	OutTopic    []string
	GroupID     string
	CallBack    CallBack
	Logger      sarama.StdLogger

	signals  chan os.Signal
	consumer *cluster.Consumer
	producer sarama.AsyncProducer
}

//NewStreaming New Streaming
func NewStreaming(seedbrokers, inTopic, outTopic []string, groupID string, config *cluster.Config, cb CallBack) (err error) {
	s := &Streaming{
		Config:      config,
		SeedBrokers: seedbrokers,
		InTopic:     inTopic,
		OutTopic:    outTopic,
		GroupID:     groupID,
		CallBack:    cb,
		signals:     make(chan os.Signal, 1),
		Logger:      sarama.Logger,
	}

	return NewWithConfig(s)
}

//NewWithConfig New With Config
func NewWithConfig(s *Streaming) (err error) {
	s.signals = make(chan os.Signal, 1)
	signal.Notify(s.signals, os.Interrupt)

	s.consumer, err = cluster.NewConsumer(s.SeedBrokers, s.GroupID, s.InTopic, s.Config)
	if err != nil {
		return err
	}

	s.producer, err = sarama.NewAsyncProducer(s.SeedBrokers, &s.Config.Config)
	if err != nil {
		return err
	}

	s.processor()

	return nil
}
