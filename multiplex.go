package streaming

func (s *Streaming) modeMultiplex() {
	defer s.consumer.Close()
	defer s.producer.AsyncClose()

	for {
		select {
		case msg, ok := <-s.consumer.Messages():
			if ok {
				producerMessage, err := s.CallBack(msg, s.OutTopic)
				if err != nil {
					s.Logger.Printf("Error: %v\n", err)
					continue
				}

				for _, m := range producerMessage {
					s.producer.Input() <- m
				}

				s.consumer.MarkOffset(msg, "")
			}

		case sign, ok := <-s.signals:
			s.Logger.Printf("Signals: %v %v\n", sign, ok)
			return
		}
	}
}
