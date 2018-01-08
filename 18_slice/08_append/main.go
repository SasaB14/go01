package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	//3 je duzina - broj elemenata na koji referise slice
	//5 je kapacitet - broj elementata u nizu ispod

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[3])
}
