package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

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

	fmt.Println(utf8.RuneCountInString(str))
}
