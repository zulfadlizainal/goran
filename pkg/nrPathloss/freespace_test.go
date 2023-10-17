package nrPathloss

import (
	"testing"
)

func TestFreeSpace(t *testing.T) {
	type args struct {
		distance float64
		freq     float64
		gaintx   float64
		gainrx   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Distance 5Km Freq 3900 MHz", args{5, 3900, 0, 0}, 118.24069222725035},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FreeSpace(tt.args.distance, tt.args.freq, tt.args.gaintx, tt.args.gainrx); got != tt.want {
				t.Errorf("FreeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
