package streaming

import "github.com/Shopify/sarama"

//CallBack Call Back
type CallBack func(*sarama.ConsumerMessage, []string) ([]*sarama.ProducerMessage, error)
