# streaming
[![Go Report Card](https://goreportcard.com/badge/github.com/flyaways/streaming?style=flat-square)](https://goreportcard.com/report/github.com/flyaways/streaming)
[![Build Status Travis](https://travis-ci.org/flyaways/streaming.svg?branch=master)](https://travis-ci.org/flyaways/streaming)
[![Build Status Semaphore](https://semaphoreci.com/api/v1/flyaways/streaming/branches/master/shields_badge.svg)](https://semaphoreci.com/flyaways/streaming)
[![LICENSE](https://img.shields.io/badge/licence-Apache%202.0-brightgreen.svg?style=flat-square)](https://github.com/flyaways/streaming/blob/master/LICENSE)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

![streaming](./kafka_diagram.png "streaming")

Streaming is a client library, where the input and output data are stored in Kafka clusters.

## Introduction

Streaming is a library written for kafka streamming processor,.

## Basic Usage

### Installation

```sh
go get -u github.com/flyaways/streaming
```

### Usage

> Streaming Processor

```go
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

```

## Credits

- [github.com/Shopify/sarama](https://github.com/Shopify/sarama)
- [github.com/bsm/sarama-cluster](https://github.com/bsm/sarama-cluster)

## Licenses

[https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fflyaways%2Fstreaming.svg?type=large)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fflyaways%2Fstreaming?ref=badge_large)
