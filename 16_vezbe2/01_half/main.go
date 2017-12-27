package main

import "fmt"

func main() {
	var x int
	fmt.Print("Unesite broj: ")
	fmt.Scan(&x)
	//fmt.Println(x)

	polutka, parnost := half(x)
	fmt.Println(polutka, parnost)
}

func half(broj int) (int, bool) {
	/* moje resenje
	  polovina := broj / 2
		return polovina, broj%2 == 0 */

	return broj / 2, broj%2 == 0

}
