package common

import "testing"

func TestDistanceInt(t *testing.T) {
	type args struct {
		a, b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,3", args{1, 3}, 2},
		{"3,1", args{3, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistanceInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
