package nrConversion

import (
	"reflect"
	"testing"

	"github.com/zulfadlizainal/goran/pkg/nrConversion"
)

func TestBandwidthToRB(t *testing.T) {
	type args struct {
		bandwidth int
		ranges    string
		scs       int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"FR1 15kHz SCS", args{50, "Sub6", 15}, 270},
		{"FR1 30kHz SCS", args{50, "Sub6", 30}, 133},
		{"FR2 15kHz SCS", args{50, "mmWave", 60}, 66},
		{"Nil", args{400, "mmWave", 60}, nil},
		{"Bandwidth (Out of Range)", args{600, "mmWave", 60}, -1},
		{"Frequency Ranges (Invalid)", args{400, "TeraHz", 120}, -1},
		{"SCS (Out of Range)", args{100, "mmWave", 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nrConversion.BandwidthToRB(tt.args.bandwidth, tt.args.ranges, tt.args.scs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BandwidthToRB() = %v, want %v", got, tt.want)
			}
		})
	}
}
