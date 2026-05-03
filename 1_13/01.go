package main

import (
	"errors"
)

var ErrDivisionByZero = errors.New("попытка деления на 0")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	return a / b, nil
}
