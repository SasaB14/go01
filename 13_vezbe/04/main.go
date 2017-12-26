package main

import "fmt"

func main() {
	var veliki int32
	var mali int32
	fmt.Print("unesite veliki broj: ")
	fmt.Scan(&veliki)
	fmt.Print("unesite mali broj: ")
	fmt.Scan(&mali)
	fmt.Println("Ostatak u deljenju velikog broja malim je:", veliki%mali)

}
