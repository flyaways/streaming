package streaming

import (
	"testing"

	cluster "github.com/bsm/sarama-cluster"
)

func TestNewStreaming(t *testing.T) {
	type args struct {
		seedbrokers []string
		inTopic     []string
		groupID     string
		outTopic    string
		config      *cluster.Config
		cb          CallBack
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewStreaming(tt.args.seedbrokers, tt.args.inTopic, tt.args.groupID, tt.args.outTopic, tt.args.config, tt.args.cb); (err != nil) != tt.wantErr {
				t.Errorf("NewStreaming() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewWithConfig(t *testing.T) {
	type args struct {
		s *Streaming
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewWithConfig(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("NewWithConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
