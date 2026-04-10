package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const vowels = "аеёиоуыэюя"

func countVowels(str string) int {
	vowelsCount := 0
	for _, v := range str {
		if strings.Contains(vowels, string(v)) {
			vowelsCount++
		}
	}

	return vowelsCount
}

func main() {
	var str string
	fmt.Print("Введите строку: ")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	str = strings.TrimSpace(str)

	fmt.Println(countVowels(str))
}
