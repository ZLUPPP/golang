package main

import (
	"fmt"
	"time"
)

func helper() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}
func a() {
	defer func() { // will NOT recover
		helper()
	}()
	panic("boom A")
}
func b() {
	defer helper() // will recover
	panic("boom B")
}
func c() {
	fmt.Println("c: recover outside defer returned", recover())
}

type Resource struct {
	Name string
}

func Open(name string) (*Resource, error) {
	fmt.Println("open resource", name)
	return &Resource{name}, nil
}

func (r *Resource) Close() {
	fmt.Println("close resource", r.Name)
}

func processAllBad(names []string) error {
	defer timer("processAllBad")()
	for _, name := range names {
		resource, err := Open(name)

		if err != nil {
			return err
		}

		defer resource.Close()
	}
	return nil
}

func process(name string) error {
	resource, err := Open(name)
	if err != nil {
		return err
	}
	// process

	defer resource.Close()
	return nil
}

func processAllGood(names []string) error {
	defer timer("processAllGood")()
	for _, name := range names {
		err := process(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Println(name, time.Since(start))
	}
}

func main() {
	files := []string{"file A", "file B", "file C"}

	err := processAllBad(files)
	if err != nil {
		fmt.Println(err)
	}

	err = processAllGood(files)
	if err != nil {
		fmt.Println(err)
	}
}
