package main

import "fmt"

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	sum := 0.0

	for _, grade := range s.Grades {
		sum += grade
	}

	return sum / float64(len(s.Grades))
}

func main() {
	me := Student{
		Name:   "Anatoly Udavikhin",
		Grades: []float64{3, 4, 5, 4, 3, 4, 5},
	}

	fmt.Println(me.AverageGrade())
}
