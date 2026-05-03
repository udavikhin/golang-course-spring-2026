package main

import "testing"

func TestSumAll(t *testing.T) {
	var testCases = []struct {
		description string
		input       []int
		want        int
	}{
		{"Positive test on 10, 20, 30, 40, 50", []int{10, 20, 30, 40, 50}, 150},
		{"Positive test on -60, -40, 50, 50", []int{-60, -40, 50, 50}, 0},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result := SumAll(tt.input...)
			if result != tt.want {
				t.Errorf("Incorrect test, want: %d, but got %d", tt.want, result)
			}
		})
	}
}
