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

func ParseAll(input string) ([]Record, error) {
	var err error = nil
	strs := strings.Split(input, "\n")
	result := make([]Record, len(strs))
	for i, line := range strs {
		line_ := strings.Split(line, ", ")
		if len(line_) != 2 {
			errors.Join(err, errors.New(fmt.Sprintf("line %d: invalid line", i)))
			continue
		}
		grade, err1 := strconv.Atoi(line_[1])
		if err1 != nil || (0 > grade || grade > 100) {
			errors.Join(err, errors.New(fmt.Sprintf("line %d: invalid grade", i)))
			continue
		}

		append(result, Record{line_[0], grade})
	}
	return result, err
}

func main() {
	input := `Aruzhan, 95
Dias, 8O
Timur
Bota, 99999999999999999999
Ali, 70`
	ParseAll(input)
}
