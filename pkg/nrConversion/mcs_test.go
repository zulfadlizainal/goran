package nrConversion

import (
	"reflect"
	"testing"
)

func TestMcsToQm(t *testing.T) {
	type args struct {
		mcs      int
		tablenum int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test Valid for Table 1", args{1, 1}, 2, false},
		{"Test Valid for Table 2", args{11, 2}, 6, false},
		{"Test Valid for Table 3", args{26, 3}, 6, false},
		{"Test Table Out of Range", args{1, 4}, nil, true},
		{"Test MCS Index Out of Range", args{32, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := McsToQm(tt.args.mcs, tt.args.tablenum)
			if (err != nil) != tt.wantErr {
				t.Errorf("McsToQm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("McsToQm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMcsToR(t *testing.T) {
	type args struct {
		mcs      int
		tablenum int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test Valid for Table 1", args{1, 1}, 0.153320313, false},
		{"Test Valid for Table 2", args{11, 2}, 0.455078125, false},
		{"Test Valid for Table 3", args{26, 3}, 0.650390625, false},
		{"Test Table Out of Range", args{1, 4}, nil, true},
		{"Test MCS Index Out of Range", args{32, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := McsToR(tt.args.mcs, tt.args.tablenum)
			if (err != nil) != tt.wantErr {
				t.Errorf("McsToR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("McsToR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMcsToSe(t *testing.T) {
	type args struct {
		mcs      int
		tablenum int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Test Valid for Table 1", args{1, 1}, 0.3066, false},
		{"Test Valid for Table 2", args{11, 2}, 2.7305, false},
		{"Test Valid for Table 3", args{26, 3}, 3.9023, false},
		{"Test Table Out of Range", args{1, 4}, nil, true},
		{"Test MCS Index Out of Range", args{32, 1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := McsToSe(tt.args.mcs, tt.args.tablenum)
			if (err != nil) != tt.wantErr {
				t.Errorf("McsToSe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("McsToSe() = %v, want %v", got, tt.want)
			}
		})
	}
}
