package main

import "fmt"

func sumAll(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(10, -2, 4, 7))
}
