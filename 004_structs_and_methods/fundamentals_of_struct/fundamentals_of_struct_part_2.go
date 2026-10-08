package main

import "fmt"

// Part 2

type Address struct {
	City string
}

type Person struct {
	Name string
	Age int
	Address Address // Composition
}

func (p Person) Greet() string { return "Hi, I am " + p.Name } 

type Employee struct {
	Person
	Address
	Company string
}

func (e Employee) Greet() string { return "HI, I am " + e.Name + " and I am working in \"" + e.Company + "\" company" }


////////////////////////////////////////////////////////////////

type Engine struct{}
func (Engine) Start() string { return "engine started" }

type Radio struct{}
func (Radio) Start() string { return "radio started" }

type Car struct {
	Engine
	Radio
}

////////////////////////////////////////////////////////////////


type Animal struct{ Name string }

func (a Animal) Sound() string { return "..." }
func (a Animal) Describe() string { return a.Name + " says " + a.Sound() }

type Dog struct{ Animal }

func (d Dog) Sound() string { return "Woof!" }

////////////////////////////////////////////////////////////////

type Logger struct{ prefix string }

func (l *Logger) Log(msg string) { fmt.Println(l.prefix + msg)}

type Service struct {
	*Logger
	Name string
}

////////////////////////////////////////////////////////////////

