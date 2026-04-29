package main

import (
	"fmt"
)

const numbersAmount = 10

func generateNumbers(ch chan int) {
	for i := 1; i <= numbersAmount; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch chan int) {
	for {
		number, ok := <-ch
		if !ok {
			fmt.Println("Канал закрыт 🚫")
			return
		}
		fmt.Println(number)
	}
}

func main() {
	ch := make(chan int, 10)

	go generateNumbers(ch)
	printNumbers(ch)
}
