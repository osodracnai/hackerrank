package hackerrank

import "testing"

func Test_birthday(t *testing.T) {
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				s: []int32{1, 2, 1, 3, 2},
				d: 3,
				m: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: []int32{4},
				d: 4,
				m: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: []int32{2, 3, 4, 4, 2, 1, 2, 5, 3, 4, 4, 3, 4, 1, 3, 5, 4, 5, 3, 1, 1, 5, 4, 3, 5, 3, 5, 3, 4, 4, 2, 4, 5, 2, 3, 2, 5, 3, 4, 2, 4, 3, 3, 4, 3, 5, 2, 5, 1, 3, 1, 4, 2, 2, 4, 3, 3, 3, 3, 4, 1, 1, 4, 3, 1, 5, 2, 5, 1, 3, 5, 4, 3, 3, 1, 5, 3, 3, 3, 4, 5, 2},
				d: 26,
				m: 8,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := birthday(tt.args.s, tt.args.d, tt.args.m); got != tt.want {
				t.Errorf("birthday() = %v, want %v", got, tt.want)
			}
		})
	}
}
