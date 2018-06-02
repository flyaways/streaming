package streaming

import (
	"os"
	"testing"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func TestStreaming_modeMultiplex(t *testing.T) {
	type fields struct {
		SeedBrokers []string
		Config      *cluster.Config
		InTopic     []string
		OutTopic    []string
		GroupID     string
		CallBack    CallBack
		Logger      sarama.StdLogger
		signals     chan os.Signal
		consumer    *cluster.Consumer
		producer    sarama.AsyncProducer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Streaming{
				SeedBrokers: tt.fields.SeedBrokers,
				Config:      tt.fields.Config,
				InTopic:     tt.fields.InTopic,
				OutTopic:    tt.fields.OutTopic,
				GroupID:     tt.fields.GroupID,
				CallBack:    tt.fields.CallBack,
				Logger:      tt.fields.Logger,
				signals:     tt.fields.signals,
				consumer:    tt.fields.consumer,
				producer:    tt.fields.producer,
			}
			s.modeMultiplex()
		})
	}
}
