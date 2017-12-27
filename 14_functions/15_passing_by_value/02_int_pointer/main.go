package main

import "fmt"

func main() {
	age := 44
	fmt.Println("Prvi print iz main funkcije")
	fmt.Println(&age) //stampa adresu promenljive age

	changeMe(&age)

	fmt.Println("Drugi print iz main funkcije")
	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println("**Prvi print iz changeMe funkcije**")
	fmt.Println(z)  //adresa pokazivaca
	fmt.Println(*z) //vrednost pokazivaca
	*z = 24
	fmt.Println("**Drugi print iz changeMe funkcije**")
	fmt.Println(z)  //adresa pokazivaca
	fmt.Println(*z) //vrednost pokazivaca
}
