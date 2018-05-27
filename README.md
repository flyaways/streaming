streaming
====================
![streaming](./kafka_diagram.png "streaming")

<!-- TOC -->

- [streaming](#streaming)
	- [Introduction](#introduction)
	- [Basic Usage](#basic-usage)
		- [Installation](#installation)
		- [Usage](#usage)
	- [Credits](#credits)
	- [Licenses](#licenses)

<!-- /TOC -->

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
