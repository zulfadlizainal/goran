package nrDownlink

import "testing"

func TestTbs(t *testing.T) {
	type args struct {
		mod       string
		nlayers   int
		nprb      int
		nreperprb int
		tcr       float64
		xoh       int
		tbscaling float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Values:", args{"16QAM", 4, 52, 120, 0.48, 6, 0.25}, 11272},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tbs(tt.args.mod, tt.args.nlayers, tt.args.nprb, tt.args.nreperprb, tt.args.tcr, tt.args.xoh, tt.args.tbscaling); got != tt.want {
				t.Errorf("Tbs() = %v, want %v", got, tt.want)
			}
		})
	}
}
