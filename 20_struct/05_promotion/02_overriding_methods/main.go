package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("Ja sam obicna osoba")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Ja sam tajni agent 00")
}

func main() {
	p1 := person{
		Name: "Jan Fleming",
		Age:  99,
	}
	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  20,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
