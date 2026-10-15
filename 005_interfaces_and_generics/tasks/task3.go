package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string { return "invalid field: " + e.Field }

func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}
func main() {
	err := validate(20)
	fmt.Println("Valid age error:", err == nil)
	fmt.Printf("%T\n", err)

	err = validate(-1)

	if err != nil {
		fmt.Println("error:", err)
	}

	var validationError *ValidationError

	if errors.As(err, &validationError) {
		fmt.Println("Field:", validationError.Field)
	}
}
