package main

import (
	"fmt"
)


func main() {
	var s1 Student // zero initialization
	s2 := Student{Name: "Aruzhan", ID: 1, GPA: 3.8} // named fields initialization ( recomendasion! )
	s3 := Student{"Dias", 2, 3.2, []string{"Go", "Databases"}} // positional initialization ( not recomended )

	s4 := &Student{Name: "Timur"} // pointer to struct
	s4.GPA = 3.5 // same as (*p).GPA = 3.5
	// q := new(Student) // pointer to zero struct

	fmt.Printf("%v\n", s1) // unnamed print
	fmt.Printf("%+v\n", s2) // named print
	fmt.Printf("%#v\n", s3) // full structure print
	fmt.Println(s1.Cources == nil)  // true

	t1 := Team{Name: "Gophers", Scores: [3]int{1, 2, 3}, Members: []string{"Ali", "Bota"}}
	t2 := t1 // copy

	t2.Name = "Rustaceans"
	t2.Scores[0] = 100
	t2.Members[0] = "Nurlan"  // slice -> link to arrey that initialized in t1

	fmt.Println(t1.Name, t1.Scores, t1.Members) // Gophers [1 2 3] [Nurlan Bota]
	fmt.Println(t2.Name, t2.Scores, t2.Members) // Rustaceans [100 2 3] [Nurlan Bota]

	a, b := Point{1, 2}, Point{1, 2}
	c, d := &Point{1, 2}, &Point{1, 2}

	fmt.Println(a == b) // val == val | true
	fmt.Println(c == d) // pointer != pointer ( different adresses 0x.. ) | true
	fmt.Println(*c == *d) // val == val | true
	fmt.Println(a == *c)  // val == val | true

	// Anonymous and empty structs
	point := struct{ X, Y int }{1, 2}

	seen := map[string]struct{}{}
	seen["go"] = struct{}{}
	z, ok := seen["go"]
	if ok {fmt.Printf("%#v\n", z)}

	fmt.Printf("%#v\n", point) // struct { X int; Y int }{X:1, Y:2}

	// Part 2

	p := Person{Name: "Ali", Address: Address{City: "Almaty"}}
	fmt.Println(p.Address.City) // Almaty


	e := Employee{
		Person: Person{Name: "Baha", Age: 22},
		Address: Address{City: "Almaty"},
		Company: "OOO tmiv",
	}
	fmt.Println(e.Name) // Baha
	fmt.Println(e.Person.Name) // Baha
	fmt.Println(e.City) // Almaty
	fmt.Println(e.Person.Greet()) // Hi, I am Baha
	fmt.Println(e.Greet()) // HI, I am Baha and I am working in "OOO tmiv" company

	car := Car{}
	fmt.Println(car.Engine.Start()) // OK
	// fmt.Println(c.Start()) // compile error


	dog := Dog{Animal{Name: "Rex"}}
	fmt.Println(dog.Sound())
	fmt.Println(dog.Animal.Sound())
	fmt.Println(dog.Describe()) // Rex says ... | ... -> cause it uses a.Sound() not d.Sound()

	ok_logger := Service{Logger: &Logger{prefix: "[svc] "}, Name: "users"}
	ok_logger.Log("started")

	// bad := Service{Name: "orders"}
	// bad.Log("started") // panic 


	// Part 3: Methods and Receivers

	r := Rectangle{3, 4}
	fmt.Println(r.Area())
	fmt.Println(Area(r))
	fmt.Println(Rectangle.Area(r))

	fmt.Println(Celsius(100).ToFahrenheit())

	// Part 4: Value vs pointer receivers

	counter := Counter{}
	counter.IncValue()
	fmt.Println(counter.n)
	counter.IncPointer()
	fmt.Println(counter.n)
	(&counter).IncPointer()
	fmt.Println(counter.n)

	accounts := []Account{{"Ali", 100}, {"Bota", 200}}
	for _, acc := range accounts { // acc = copy or each element
		acc.Deposit(50)
	}

	fmt.Println(accounts) // [{Ali 100} {Bota 200}]

	for i := range accounts {
		accounts[i].Deposit(50)
	}

	fmt.Println(accounts) // [{Ali 150} {Bota 250}]

	// Part 5: Method sets

	var _ Speaker = &ByValue{} // *ByValue | Hello + Speak
	// var _ Speaker = ByValue{} // Hello | error
	var _ Speaker = ByPointer{Base: &Base{}} // ByPoiner | Hello + Speak
	var _ Speaker = &ByPointer{Base: &Base{}} // *ByPoiner | Hello + Speak


	book := Book{Title: "The Go Programming Language", Author: "Donovan & Kernighan"}
	fmt.Println(book)
	fmt.Println(&book)


}
