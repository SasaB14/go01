package main

import "fmt"

func main() {
	var x int
	fmt.Print("Unesite broj: ")
	fmt.Scan(&x)

	rezultat := func(broj int) (int, bool) {
		return broj / 2, broj%2 == 0
	}
	fmt.Println(rezultat(x))
}
