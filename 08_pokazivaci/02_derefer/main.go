package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  //stampa vrednost
	fmt.Println(&a) //stampa adresu

	var b *int = &a

	fmt.Println(b)  //stampa adresu
	fmt.Println(*b) //stampa vrednost

	//b je int pokazivac
	//b pokazuje na memorijsku adresu gde  se cuva int
	//da bi videli vrednost koja se cuva na toj adresi dodajemo * ispred pokazivaca
	//ovo je dereferencing
	//* je operator u ovom slucaju
}
