package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d\n", b.Title, b.Author, b.Year)
}

func main() {
	metro := Book{
		Title:  "Metro 2033",
		Author: "Dmitry Glukhovsky",
		Year:   2002,
	}

	fmt.Println(metro.GetInfo())
}
