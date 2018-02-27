package streaming

import (
	cluster "github.com/bsm/sarama-cluster"
)

func (s *Streaming) modePartitions() {
	defer s.consumer.Close()
	defer s.producer.AsyncClose()

	for {
		select {
		case part, ok := <-s.consumer.Partitions():
			if ok {
				go func(pc cluster.PartitionConsumer) {

					for msg := range pc.Messages() {

						producerMessage, err := s.CallBack(msg, s.OutTopic)
						if err != nil {
							s.Logger.Printf("Error: %v\n", err)
						}

						s.producer.Input() <- producerMessage
						s.consumer.MarkOffset(msg, "")
					}
				}(part)
			}

		case sign, ok := <-s.signals:
			s.Logger.Printf("Signals: %v %v\n", sign, ok)
			return
		}
	}
}
