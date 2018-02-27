package streaming

import (
	cluster "github.com/bsm/sarama-cluster"
)

func (s *Streaming) processor() {
	if s.Config.Consumer.Return.Errors {
		go func() {
			for err := range s.consumer.Errors() {
				s.Logger.Printf("Error: %v\n", err)
			}
		}()
	}

	if s.Config.Group.Return.Notifications {
		go func() {
			for ntf := range s.consumer.Notifications() {
				s.Logger.Printf("Notification: %+v\n", ntf)
			}
		}()
	}

	if s.Config.Group.Mode == cluster.ConsumerModePartitions {
		go s.modePartitions()
	}

	if s.Config.Group.Mode == cluster.ConsumerModeMultiplex {
		go s.modeMultiplex()
	}
}
