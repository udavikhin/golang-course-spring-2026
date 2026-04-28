package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWordEntries(str string) map[string]int {
	parsedString := strings.Split(str, " ")

	result := make(map[string]int)

	for _, word := range parsedString {
		if len(word) == 0 {
			continue
		}

		result[word]++
	}

	return result
}

func prettyPrintMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%-30s: %d\n", key, value)
	}
}

func main() {
	fmt.Print("Введите строку: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	str = strings.TrimSpace(str)
	wordEntriesCount := countWordEntries(str)
	prettyPrintMap(wordEntriesCount)
}
