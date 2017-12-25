package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Unesite preplivane metre:")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " metara su ", yards, " jardi.")
}
