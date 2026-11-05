package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("Not Found")

type GradeError struct {
	Name  string
	Grade int
}

func (e GradeError) Error() string {
	return fmt.Sprintf("grade %d is out of range from 0 to 100", e.Grade)
}

type GradeBook map[string]int

func (g GradeBook) Get(name string) (int, error) {
	val, ok := g[name]
	if !ok {
		return 0, ErrNotFound
	}

	return val, nil
}

func (g GradeBook) Set(name string, grade int) error {
	if 0 > grade || grade > 100 {
		return GradeError{Name: name, Grade: grade}
	}
	g[name] = grade
	return nil
}

func main() {
	ge := make(GradeBook)
	ge.Set("Alice", 88)
	value, err := ge.Get("Bob")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	err = ge.Set("Bob", 101)
	if err != nil {
		fmt.Println(err)
	}
}
