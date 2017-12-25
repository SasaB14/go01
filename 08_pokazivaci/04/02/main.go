package main

import "fmt"

func zero(z int) {
	fmt.Printf("adresa u funkciji zero %p\n", &z) //adresa u funkciji zero
	fmt.Println(&z)                               //adresa u funkciji zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("adresa u main funkciji %p\n", &x) //adresa u main funkciji
	fmt.Println(&x)                               //adresa u main funkciji
	zero(x)
	fmt.Println(x) //x je i dalje 5
}
