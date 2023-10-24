package nrMeasurement

import "testing"

func TestSsRsrp(t *testing.T) {
	type args struct {
		sssrepow float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Convert ", args{0.0000007589472}, -84.99999678657063},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SsRsrp(tt.args.sssrepow); got != tt.want {
				t.Errorf("SsRsrp() = %v, want %v", got, tt.want)
			}
		})
	}
}
