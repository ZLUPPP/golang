package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Record struct {
	Name  string
	Grade int
}

var ErrBadFormat = errors.New("bad format")

func ParseAll(input string) ([]Record, error) {
	var err error = nil
	strs := strings.Split(input, "\n")
	var result []Record

	for i, line := range strs {
		line_ := strings.Split(line, ",")
		if len(line_) != 2 {
			err = errors.Join(err, fmt.Errorf("line %d: %w", i+1, ErrBadFormat))
			continue
		}
		grade, err1 := strconv.Atoi(strings.TrimSpace(line_[1]))
		if err1 != nil {
			err = errors.Join(err, fmt.Errorf("line %d: %w", i+1, err1))
			continue
		}

		if 0 > grade || grade > 100 {
			err = errors.Join(err, fmt.Errorf("line %d: grade out of range: %d", i+1, grade))
			continue
		}

		result = append(result, Record{line_[0], grade})
	}
	return result, err
}

func main() {
	input := `Aruzhan, 95
Dias, 8O
Timur
Bota, 99999999999999999999
Ali, 70`
	records, err := ParseAll(input)

	fmt.Println("Records: ", records)
	fmt.Println("Errors: ", err)

	fmt.Println("Bad format: ", errors.Is(err, ErrBadFormat))
	fmt.Println("Syntax: ", errors.Is(err, strconv.ErrSyntax))
	fmt.Println("Range: ", errors.Is(err, strconv.ErrRange))

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Println("NumError: ", numErr)
	}
}
