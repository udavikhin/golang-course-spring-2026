package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("попытка деления на 0")
	}

	return a / b, nil
}

func main() {
	res, err := divide(10.6, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
