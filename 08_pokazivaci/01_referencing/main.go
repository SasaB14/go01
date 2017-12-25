package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)

	//kod iznad cini b pokazivacem na memorijsku adresu gde se int cuva
	//b je tipa pokazivac
	//*int -- * je deo tipa - b je tipa *int
}
