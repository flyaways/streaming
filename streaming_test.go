package streaming

import "testing"

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
