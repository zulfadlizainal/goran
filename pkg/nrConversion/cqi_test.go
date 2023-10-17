package nrConversion

import (
	"reflect"
	"testing"
)

func TestQCIToType(t *testing.T) {
	type args struct {
		qci int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"5QI 1", args{1}, "GBR", false},
		{"5QI 2", args{2}, "GBR", false},
		{"5QI 3", args{3}, "GBR", false},
		{"5QI -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QCIToType(tt.args.qci)
			if (err != nil) != tt.wantErr {
				t.Errorf("QCIToType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"5QI 1", args{1}, 2.0, false},
		{"5QI 2", args{2}, 4.0, false},
		{"5QI 3", args{3}, 3.0, false},
		{"5QI -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QCIToPriority(tt.args.qci)
			if (err != nil) != tt.wantErr {
				t.Errorf("QCIToPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"5QI 1", args{1}, 100, false},
		{"5QI 2", args{2}, 150, false},
		{"5QI 3", args{3}, 50, false},
		{"5QI -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QCIToPacketDelay(tt.args.qci)
			if (err != nil) != tt.wantErr {
				t.Errorf("QCIToPacketDelay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"5QI 1", args{1}, 0.01, false},
		{"5QI 2", args{2}, 0.001, false},
		{"5QI 3", args{3}, 0.001, false},
		{"5QI -1 (Out of Range)", args{-1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QCIToPacketLoss(tt.args.qci)
			if (err != nil) != tt.wantErr {
				t.Errorf("QCIToPacketLoss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QCIToPacketLoss() = %v, want %v", got, tt.want)
			}
		})
	}
}
