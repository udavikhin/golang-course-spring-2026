package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

const arrLength = 10

func fillArrayWithRandomValues(arr [arrLength]uint) [arrLength]uint {
	for i, _ := range arr {
		arr[i] = rand.UintN(100) + 1
	}

	return arr
}

func main() {
	var arr [arrLength]uint

	arr = fillArrayWithRandomValues(arr)

	slice := make([]uint, arrLength)
	copy(slice, arr[:])
	slices.Sort(slice)

	fmt.Println("Исходный массив: ", arr)
	fmt.Println("Отсортированный слайс: ", slice)
}
