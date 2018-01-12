package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			first: "Miss",
			last:  "Moneypenny",
			age:   19,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.first, p1.last, p1.age, p1.LicenseToKill)
	fmt.Println(p2.first, p2.last, p2.age, p2.LicenseToKill)
}
