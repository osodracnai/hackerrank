package hackerrank

import "testing"

func Test_flippingMatrix(t *testing.T) {
	type args struct {
		matrix [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				matrix: [][]int32{
					{112, 42, 83, 119},
					{56, 125, 56, 49},
					{15, 78, 101, 43},
					{62, 98, 114, 108},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippingMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("flippingMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
