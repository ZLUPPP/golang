package main

import "fmt"

type Student struct {
	Name string
	ID int
	GPA float64
	Cources []string
}

func main() {
	var s1 Student // zero initialization
	s2 := Student{Name: "Aruzhan", ID: 1, GPA: 3.8} // named fields initialization ( recomendasion! )
	s3 := Student{"Dias", 2, 3.2, []string{"Go", "Databases"}} // positional initialization ( not recomended )

	s4 := &Student{Name: "Timur"} // pointer to struct
	s4.GPA = 3.5 // same as (*p).GPA = 3.5

	// Task 1: Student records

	fmt.Printf("%+v\n", s1) // {Name: ID:0 GPA:0 Cources:[]}
	fmt.Printf("%+v\n", s2) // {Name:Aruzhan ID:1 GPA:3.8 Cources:[]}
	fmt.Printf("%+v\n", s3) // {Name:Dias ID:2 GPA:3.2 Cources:[Go Databases]}
	fmt.Printf("%+v\n", s4) // &{Name:Timur ID:0 GPA:3.5 Cources:[]}

	addCourse(&s1, "Aurafarming")
	fmt.Printf("%+v\n", s1) // Nothing change because when we throw struct to func it just copy itselt into function input
}

func addCourse(s *Student, course string) {
	s.Cources = append(s.Cources, course)
}
