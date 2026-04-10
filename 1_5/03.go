package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalizeWords(s string) string {
	parts := strings.Split(s, " ")
	var capitalizedString []string
	for _, word := range parts {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])

		capitalizedString = append(capitalizedString, string(runes))
	}

	return strings.Join(capitalizedString, " ")
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
	str = capitalizeWords(str)

	fmt.Println(str)
}
