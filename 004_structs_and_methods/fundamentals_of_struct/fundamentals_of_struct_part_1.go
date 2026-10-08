package main

// Part 1

type Student struct {
	Name string
	ID int
	GPA float64
	Cources []string
}

type Team struct {
	Name string
	Scores [3]int
	Members []string
}

type Point struct { X, Y int }

func addCourse(s *Student, course string) {
	s.Cources = append(s.Cources, course)
}
