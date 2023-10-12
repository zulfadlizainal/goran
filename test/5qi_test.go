package nrConversion

import (
	"testing"

	"github.com/zulfadlizainal/goran/pkg/nr/nrConversion"
)

func TestQCIToType(t *testing.T) {
	type args struct {
		qci int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"5QI 1", args{1}, "GBR"},
		{"5QI 2", args{2}, "GBR"},
		{"5QI 3", args{3}, "GBR"},
		{"5QI -1 (Out of Range)", args{-1}, "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.QCIToType(tt.args.qci); got != tt.want {
				t.Errorf("QCIToType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQCIToPriority(t *testing.T) {
	type args struct {
		qci int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"5QI 1", args{1}, 2.0},
		{"5QI 2", args{2}, 4.0},
		{"5QI 3", args{3}, 3.0},
		{"5QI -1 (Out of Range)", args{-1}, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.QCIToPriority(tt.args.qci); got != tt.want {
				t.Errorf("QCIToPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQCIToPacketDelay(t *testing.T) {
	type args struct {
		qci int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5QI 1", args{1}, 100},
		{"5QI 2", args{2}, 150},
		{"5QI 3", args{3}, 50},
		{"5QI -1 (Out of Range)", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.QCIToPacketDelay(tt.args.qci); got != tt.want {
				t.Errorf("QCIToPacketDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQCIToPacketLoss(t *testing.T) {
	type args struct {
		qci int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"5QI 1", args{1}, 0.01},
		{"5QI 2", args{2}, 0.001},
		{"5QI 3", args{3}, 0.001},
		{"5QI -1 (Out of Range)", args{-1}, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.QCIToPacketLoss(tt.args.qci); got != tt.want {
				t.Errorf("QCIToPacketLoss() = %v, want %v", got, tt.want)
			}
		})
	}
}
