package main

import "fmt"

func zero(z *int) {
	fmt.Println("adresa z iz zero:")
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("adresa x iz main:")
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
