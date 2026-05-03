package main

import (
	"errors"
	"testing"
)

type divideOutput struct {
	value float64
	err   error
}

func TestDivide(t *testing.T) {
	var testCases = []struct {
		description string
		input       [2]float64
		want        divideOutput
	}{
		{"Positive test on 16 and 2", [2]float64{16, 2}, divideOutput{8, nil}},
		{"Positive test on 111 and 3", [2]float64{111, 3}, divideOutput{37, nil}},
		{"Positive test on 20.5 and 5", [2]float64{20.5, 5}, divideOutput{4.1, nil}},
		{"Positive test on 10 and 0.1", [2]float64{10, 0.1}, divideOutput{100, nil}},
		{"Negative test on 0", [2]float64{1, 0}, divideOutput{0, ErrDivisionByZero}},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, err := Divide(tt.input[0], tt.input[1])
			if result != tt.want.value || !errors.Is(err, tt.want.err) {
				t.Errorf("Incorrect test, want: (%.2f, %s), but got %.2f", tt.want.value, tt.want.err, result)
			}
		})
	}
}
