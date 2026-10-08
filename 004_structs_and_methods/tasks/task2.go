package main

import "fmt"

// Task 2: University

type Person struct{
	Name string
	Email string
}

func (p Person) Contact() string { return p.Name + " <" + p.Email + ">" }

type Student struct{
	Person
	GPA float64
} 

type Teacher struct{
	Person
	Department string
}

func (t Teacher) Contact() string { return t.Name + " <" + t.Email + "> | " + t.Department }


type WID struct{
	ID int
}

type SID struct{
	ID int
}

type WEED struct{
	WID
	SID
}


func main() {
	teacher := Teacher{Person: Person{Name: "Egor", Email: "cosogor@gmail.com"}, Department: "SAM"}


	fmt.Println(teacher.Contact())
	fmt.Println(teacher.Person.Contact())

	weed := WEED{WID{ID: 1}, SID{ID: 2}}

	// fmt.Println(weed.ID) // ambiguous selector weed.ID
	fmt.Println(weed.WID.ID)


}
