package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "世界"
	fmt.Println(len(s))                    // 6 bytes = 2 symbols ( rune ) * 3 bytes
	fmt.Println(utf8.RuneCountInString(s)) // 2 symbols ( rune )

	for i := range s {
		fmt.Printf("%d: %c\n", i, s[i]) // Here we iterate through []byte(s)
	}

	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r) // Here we iterate through []rune(s)
	}

	// How can you access the second rune in a string in a readable form? And how does it work?

	fmt.Printf("%c", []rune(s)[1]) // Here we access the second rune
	// we convert slice of bytes to slice of runes and access second rune

	// Trim functions

	fmt.Println(strings.TrimRight("123oxo", "ox"))  // 123
	fmt.Println(strings.TrimLeft("123oxo", "321"))  // oxo
	fmt.Println(strings.TrimSuffix("123oxo", "xo")) // 123o
	fmt.Println(strings.TrimPrefix("123oxo", "12")) // 3oxo
	fmt.Println(strings.TrimSpace("    123oxo  "))  // 123oxo

	fmt.Println(concatWrong([]string{"abc", "efg"}))
	fmt.Println(concatRight([]string{"abc", "efg"}))

}

// bull shift strings concatenation 👎
func concatWrong(values []string) string {
	s := ""
	for _, v := range values {
		s += v + "\n" // allocate memory for new string every iteration and copy all elements from previous string
	}
	return s
}

// perfect string concatenation 👍
func concatRight(values []string) string {
	var sb strings.Builder
	length := 0

	for _, v := range values {
		length += len(v) + 1
	}

	sb.Grow(length) // Extend size to needed

	for _, v := range values {
		sb.WriteString(v)
		sb.WriteByte('\n') // write in buffer of string builder
	}
	return sb.String()
}

// Wrong function with unnecessary type conversions
// 65.56 ns/op   32 B/op  2 allocs/op
func wrongProcessPayload(data []byte) string {
	trimmed := string(bytes.TrimSpace(data))       // allocate new memory for string ([]byte -> string)
	clean := strings.ReplaceAll(trimmed, "\r", "") // strings are immutable -> allocates new memory
	return clean
}

// Correct function
// 44.30 ns/op   16 B/op  1 allocs/op
// ~1.48 time more efficient
func correctProcessPayload(data []byte) []byte {
	trimmed := bytes.TrimSpace(data)                             // do not allocate memory, return slice of bytes array
	clean := bytes.ReplaceAll(trimmed, []byte("\r"), []byte("")) // allocates memory for result when replacement needed
	return clean
}
