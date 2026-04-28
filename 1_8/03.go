package main

import "fmt"

func main() {
	value := 0

	counter := func() int {
		value++
		return value
	}

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
