package main

type User struct {
	Name string
	Age  int
}

func newUserValue(name string) User {
	u := User{Name: name, Age: 20}
	return u
}

func newUserPointer(name string) *User {
	u := User{Name: name, Age: 20}
	return &u
}

func f1() int {
	x := 42
	return x
}

func f2() *int {
	x := 42
	return &x
}

func main() {

}
