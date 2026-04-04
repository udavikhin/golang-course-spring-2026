package main

import "fmt"

const numbersCount = 5

func sortDesc(arr [numbersCount]uint) [numbersCount]uint {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}

func calcArithmeticMean(arr [numbersCount]uint) float32 {
	sum := uint(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return float32(sum) / numbersCount
}

func main() {
	var numbers [numbersCount]uint
	fmt.Printf("Введите %d натуральных чисел: \n", numbersCount)
	for i := 0; i < numbersCount; i++ {
		fmt.Scanln(&numbers[i])
	}
	fmt.Println(numbers)
	fmt.Println("Отсортированные элементы: ", sortDesc(numbers))
	fmt.Println("Самое большое число: ", sortDesc(numbers)[0])
	fmt.Println("Самое маленькое число: ", sortDesc(numbers)[numbersCount-1])
	fmt.Println("Среднее арифметическое: ", calcArithmeticMean(numbers))
}
