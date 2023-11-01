package nrMeasurement

import "testing"

func TestSsRssi(t *testing.T) {
	type args struct {
		ss_rsrp float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test valid:", args{-85}, -61.82981898951888},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SsRssi(tt.args.ss_rsrp); got != tt.want {
				t.Errorf("SsRssi() = %v, want %v", got, tt.want)
			}
		})
	}
}
