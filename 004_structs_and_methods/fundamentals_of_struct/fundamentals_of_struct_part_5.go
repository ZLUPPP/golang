package main

// Part 5: Method sets

type Speaker interface{ Speak() string }

type Base struct{}
func (Base) Hello() string { return "hello" }
func (*Base) Speak() string { return "speak" }

type ByValue struct{ Base } // embeds Base

type ByPointer struct{ *Base } // embeds *Base

////////////////////////////////////////////////////////////////

type Book struct {
	Title string
	Author string
}

func (b *Book) String() string { return b.Title + " by " + b.Author }