package nrConversion

import (
	"testing"

	"github.com/zulfadlizainal/goran/pkg/nrConversion"
)

func TestNumerologyToScs(t *testing.T) {
	type args struct {
		numerology int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Numerology 0", args{0}, 15},
		{"Numerology 1", args{1}, 30},
		{"Numerology 2", args{2}, 60},
		{"Numerology 3", args{3}, 120},
		{"Numerology 4", args{4}, 240},
		{"Numerology -1 (Out of Range)", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.NumerologyToScs(tt.args.numerology); got != tt.want {
				t.Errorf("NumerologyToScs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumerologyToSymbolPerSlot(t *testing.T) {
	type args struct {
		numerology int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Numerology 0", args{0}, 14},
		{"Numerology 1", args{1}, 14},
		{"Numerology 2", args{2}, 14},
		{"Numerology 3", args{3}, 14},
		{"Numerology 4", args{4}, 14},
		{"Numerology 4", args{5}, 14},
		{"Numerology 4", args{6}, 14},
		{"Numerology -1 (Out of Range)", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.NumerologyToSymbolPerSlot(tt.args.numerology); got != tt.want {
				t.Errorf("NumerologyToSymbolPerSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumerologyToSlotPerSubframe(t *testing.T) {
	type args struct {
		numerology int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Numerology 0", args{0}, 1},
		{"Numerology 1", args{1}, 2},
		{"Numerology 2", args{2}, 4},
		{"Numerology 3", args{3}, 8},
		{"Numerology 4", args{4}, 16},
		{"Numerology 4", args{5}, 32},
		{"Numerology 4", args{6}, 64},
		{"Numerology -1 (Out of Range)", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.NumerologyToSlotPerSubframe(tt.args.numerology); got != tt.want {
				t.Errorf("NumerologyToSlotPerSubframe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumerologyToSlotPerFrame(t *testing.T) {
	type args struct {
		numerology int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Numerology 0", args{0}, 10},
		{"Numerology 1", args{1}, 20},
		{"Numerology 2", args{2}, 40},
		{"Numerology 3", args{3}, 80},
		{"Numerology 4", args{4}, 160},
		{"Numerology 4", args{5}, 320},
		{"Numerology 4", args{6}, 640},
		{"Numerology -1 (Out of Range)", args{-1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.NumerologyToSlotPerFrame(tt.args.numerology); got != tt.want {
				t.Errorf("NumerologyToSlotPerFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}
