package nrMeasurement

import "testing"

func TestSsSinr(t *testing.T) {
	type args struct {
		ss_rsrp       float64
		ncell_ss_rsrp float64
		scs           int
		noisefig      float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test for SCS 30kHz with 7dB noise figure", args{-80, -90, 30, 7}, 9.997401183480505},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SsSinr(tt.args.ss_rsrp, tt.args.ncell_ss_rsrp, tt.args.scs, tt.args.noisefig); got != tt.want {
				t.Errorf("SsSinr() = %v, want %v", got, tt.want)
			}
		})
	}
}
