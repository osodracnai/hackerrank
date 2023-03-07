package hackerrank

import "testing"

func Test_lonelyinteger2(t *testing.T) {
	type args struct {
		a []int32
	}
	var tests = []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				a: []int32{1, 1, 2, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lonelyinteger2(tt.args.a); got != tt.want {
				t.Errorf("lonelyinteger2() = %v, want %v", got, tt.want)
			}
		})
	}
}
