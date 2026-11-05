package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Name string
	Run  func() error
}

func runOne(task Task) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf(
				"task %s: panic: %v",
				task.Name,
				r,
			)
		}
	}()

	e = task.Run()
	if e != nil {
		return fmt.Errorf("task %s: %w", task.Name, e)
	}
	return nil
}

func SafeRun(tasks []Task) []error {
	errs := []error{}

	for _, task := range tasks {
		err := runOne(task)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

var ErrSome = errors.New("something went wrong")

func main() {
	tasks := []Task{
		{
			Name: "success",
			Run: func() error {
				return nil
			},
		},
		{
			Name: "indexOutOfRange",
			Run: func() error {
				a := []int{1, 2, 3}
				a[10] = 0
				return nil
			},
		},
		{
			Name: "writeInNilMap",
			Run: func() error {
				a := map[int]int{}
				a[1] = 0
				return nil
			},
		},
		{
			Name: "panic",
			Run: func() error {
				panic("config is missing")
			},
		},
		{
			Name: "someError",
			Run: func() error {
				return ErrSome
			},
		},
	}

	errs := SafeRun(tasks)
	err := errors.Join(errs...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("SomeError:", errors.Is(err, ErrSome))

}
