package tDirectory

import "testing"

func TestMul(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero zero pair", args{[]int{0, 0}}, 0},
		{"zero one pair", args{[]int{0, 1}}, 0},
		{"one one pair", args{[]int{1, 1}}, 1},
		{"one two pair", args{[]int{1, 2}}, 2},
		{"one two three", args{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.set...); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero zero pair", args{[]int{0, 0}}, 0},
		{"zero one pair", args{[]int{0, 1}}, 1},
		{"one one pair", args{[]int{1, 1}}, 2},
		{"one two pair", args{[]int{1, 2}}, 3},
		{"one two three", args{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.set...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
