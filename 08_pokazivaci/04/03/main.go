package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)       //prosledjujemo adresu promenljive -- funkcija zero ima argument pokazivac
	fmt.Println(x) //x je nula
}
