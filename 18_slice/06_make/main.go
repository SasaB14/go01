package main

import "fmt"

func main() {
	customerNumber := make([]int, 3)
	//3 je duzina i kapacitet
	// duzina broj elemenata na koji referise slice
	// kapacitet broj elemenata u nizu ispod slice-a
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	//3 je duzina - broj elemenata na koji referise slice
	//5 je kapacitet - broj elementata u nizu ispod

	greeting[0] = "Dobro jutro"
	greeting[1] = "Dobar dan"
	greeting[2] = "Dobro vece"

	fmt.Println(greeting[2])
}
