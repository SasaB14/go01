package main

import "fmt"

/*
func main() {
	var x int
	fmt.Print("Unesite broj: ")
	fmt.Scan(&x)

	polutka, parnost := half(x)
	fmt.Println(polutka, parnost)
}

func half(broj int) (int, bool) {
	return broj / 2, broj%2 == 0
}*/

func main() {
	var x int
	fmt.Print("Unesite broj: ")
	fmt.Scan(&x)

	aa := func(broj int) (int, bool) {
		return broj / 2, broj%2 == 0
	}
	fmt.Println(aa(x))
}
