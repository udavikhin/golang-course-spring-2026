package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validateParentheses(formula string) (bool, int, int) {
	isFormulaValid := true
	leftParenCount := 0
	rightParenCount := 0

	for _, char := range formula {
		if string(char) == "(" {
			leftParenCount++
		}
		if string(char) == ")" {
			rightParenCount++
			if rightParenCount > leftParenCount {
				isFormulaValid = false
			}
		}
	}

	if leftParenCount != rightParenCount {
		isFormulaValid = false
	}

	return isFormulaValid, leftParenCount, rightParenCount
}

func formatResultString(isFormulaValid bool, leftParenCount int, rightParenCount int) string {
	resultString := "Скобки расставлены "
	if isFormulaValid {
		resultString += "верно"
	} else {
		resultString += "неверно"
	}
	resultString += fmt.Sprintf(", %d открывающиеся, %d закрывающиеся", leftParenCount, rightParenCount)

	return resultString
}

func main() {
	var formula string
	fmt.Print("Введите формулу: ")

	reader := bufio.NewReader(os.Stdin)
	formula, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	formula = strings.TrimSpace(formula)
	result := formatResultString(validateParentheses(formula))

	fmt.Println(result)
}
