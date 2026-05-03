package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("logfile.log", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	totalErrorStrings := 0

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Достигнут конец файла")
				break
			}
		}
		if strings.Contains(line, "error") {
			totalErrorStrings++
		}
	}

	fmt.Println("Найдено строк со словом \"error\":", totalErrorStrings)
}
