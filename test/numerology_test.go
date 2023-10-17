package nrConversion

import (
	"reflect"
	"testing"

	"github.com/zulfadlizainal/goran/pkg/nrConversion"
)

func TestNumerologyToScs(t *testing.T) {
	type args struct {
		numerology int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Numerology 0", args{0}, 15, false},
		{"Numerology 1", args{1}, 30, false},
		{"Numerology 2", args{2}, 60, false},
		{"Numerology 3", args{3}, 120, false},
		{"Numerology 4", args{4}, 240, false},
		{"Numerology -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nrConversion.NumerologyToScs(tt.args.numerology)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumerologyToScs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Numerology 0", args{0}, 14, false},
		{"Numerology 1", args{1}, 14, false},
		{"Numerology 2", args{2}, 14, false},
		{"Numerology 3", args{3}, 14, false},
		{"Numerology 4", args{4}, 14, false},
		{"Numerology 4", args{5}, 14, false},
		{"Numerology 4", args{6}, 14, false},
		{"Numerology -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nrConversion.NumerologyToSymbolPerSlot(tt.args.numerology)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumerologyToSymbolPerSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Numerology 0", args{0}, 1, false},
		{"Numerology 1", args{1}, 2, false},
		{"Numerology 2", args{2}, 4, false},
		{"Numerology 3", args{3}, 8, false},
		{"Numerology 4", args{4}, 16, false},
		{"Numerology 4", args{5}, 32, false},
		{"Numerology 4", args{6}, 64, false},
		{"Numerology -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nrConversion.NumerologyToSlotPerSubframe(tt.args.numerology)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumerologyToSlotPerSubframe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Numerology 0", args{0}, 10, false},
		{"Numerology 1", args{1}, 20, false},
		{"Numerology 2", args{2}, 40, false},
		{"Numerology 3", args{3}, 80, false},
		{"Numerology 4", args{4}, 160, false},
		{"Numerology 4", args{5}, 320, false},
		{"Numerology 4", args{6}, 640, false},
		{"Numerology -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nrConversion.NumerologyToSlotPerFrame(tt.args.numerology)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumerologyToSlotPerFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumerologyToSlotPerFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}
