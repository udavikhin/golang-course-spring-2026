package main

import "fmt"

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(applyOperation(20, 30, sum))
	fmt.Println(applyOperation(100, 40, subtract))
	fmt.Println(applyOperation(60, 2, multiply))
}
