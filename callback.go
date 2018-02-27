package streaming

import "github.com/Shopify/sarama"

type CallBack func(*sarama.ConsumerMessage, string) (*sarama.ProducerMessage, error)
