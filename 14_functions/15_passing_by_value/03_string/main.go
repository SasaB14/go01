package main

import "fmt"

func main() {
	name := "sasa"
	fmt.Println(name) //sasa

	changeMe(name)

	fmt.Println(name) //sasa
}

func changeMe(z string) {
	fmt.Println(z) //sasa
	z = "Rocky"
	fmt.Println(z) //Rocky
}
