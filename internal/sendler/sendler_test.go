package sendler

import (
	"testing"
)

func TestSendServer(t *testing.T) {
	type args struct {
		site    string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendServer(tt.args.site, tt.args.message)
		})
	}
}

func TestSendTg(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendTg(tt.args.message)
		})
	}
}
