package main

import "fmt"

func main() {
	name := "sasa"
	fmt.Println(&name) //adresa promenljive name
	fmt.Println(name)

	changeMe(&name)

	fmt.Println(&name) //adresa promenljive name
	fmt.Println(name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println("*******")
	fmt.Println(z)  //adresa z
	fmt.Println(*z) //sasa
	*z = "Rocky"
	fmt.Println(z)  //adresa z
	fmt.Println(*z) //sasa
	fmt.Println("*******")
}
