package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet()) //mora greet() jer je return iz func makeGreeter() func() string
	fmt.Printf("%T \n", greet)
}
