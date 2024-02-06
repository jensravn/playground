package main

import "testing"

func Test_getPage(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Question 1, Page 1", args{1}, 1},
		{"Question 5, Page 1", args{5}, 1},
		{"Question 6, Page 2", args{6}, 2},
		{"Question 10, Page 2", args{10}, 2},
		{"Question 11, Page 3", args{11}, 3},
		{"Question 15, Page 3", args{15}, 3},
		{"Question 16, Page 4", args{16}, 4},
		{"Question 20, Page 4", args{20}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPage(tt.args.n); got != tt.want {
				t.Errorf("getPageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
