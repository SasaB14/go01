package main

import "fmt"

func main() {
	mySlice := []string{"Ponedeljak", "Utorak"}
	myOtherSlice := []string{"Sreda", "Cetvrtak", "Petak"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}
