package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("memorijska adresa za a - ", &a)
	fmt.Printf("%d \n", &a)
}
