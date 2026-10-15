package main

import (
	"errors"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%.1f°C", float64(c)) }

func describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "<nil>"
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("integer: %v", x)
	case string:
		return fmt.Sprintf("string of %d bytes", len(x))
	case []int:
		return fmt.Sprintf("slice of %d ints", len(x))
	case bool:
		return fmt.Sprintf("boolean: %v", x)
	case map[string]int:
		return fmt.Sprintf("slice of map[string]int: %v", x)
	case error:
		return "error " + x.Error()
	case fmt.Stringer:
		return "Stringer " + x.String()
	default:
		return fmt.Sprintf("unknown type: %T", x)
	}
}

func main() {
	fmt.Println(describe(nil))                        // <nil>
	fmt.Println(describe(7))                          // integer: 7
	fmt.Println(describe(int64(7)))                   // integer: 7
	fmt.Println(describe("привет"))                   // slice of 12 bytes
	fmt.Println(describe([]int{1, 2, 3}))             // slice of 3 ints
	fmt.Println(describe(errors.New("disk full")))    // error disk full
	fmt.Println(describe(Celsius(21.5)))              // Stringer 12.5C
	fmt.Println(describe(3.14))                       // unknown type: float64
	fmt.Println(describe(struct{}{}))                 // unknown type: struct {}
	fmt.Println(describe(true))                       // boolean: true
	fmt.Println(describe(map[string]int{"hello": 2})) // slice of map[string]int: map[hello:2]
}
