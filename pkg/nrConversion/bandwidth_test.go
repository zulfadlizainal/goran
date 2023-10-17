package nrConversion

import (
	"reflect"
	"testing"
)

func TestBandwidthToRB(t *testing.T) {
	type args struct {
		bandwidth int
		fr        string
		scs       int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"FR1 15kHz SCS", args{50, "Sub6", 15}, 270, false},
		{"FR1 30kHz SCS", args{50, "Sub6", 30}, 133, false},
		{"FR2 15kHz SCS", args{50, "mmWave", 60}, 66, false},
		{"Nil", args{400, "mmWave", 60}, nil, false},
		{"Bandwidth (Out of Range)", args{600, "mmWave", 60}, nil, true},
		{"Frequency Ranges (Invalid)", args{400, "TeraHz", 120}, nil, true},
		{"SCS (Out of Range)", args{100, "mmWave", 0}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BandwidthToRB(tt.args.bandwidth, tt.args.fr, tt.args.scs)
			if (err != nil) != tt.wantErr {
				t.Errorf("BandwidthToRB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BandwidthToRB() = %v, want %v", got, tt.want)
			}
		})
	}
}
