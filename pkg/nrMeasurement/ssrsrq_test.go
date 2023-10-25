package nrMeasurement

import "testing"

func TestSsRsrq(t *testing.T) {
	type args struct {
		ss_rsrp float64
		rssi    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test with 30kHz RSSI values", args{-85, -61.8298164379}, -10.159883605460191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SsRsrq(tt.args.ss_rsrp, tt.args.rssi); got != tt.want {
				t.Errorf("SsRsrq() = %v, want %v", got, tt.want)
			}
		})
	}
}
