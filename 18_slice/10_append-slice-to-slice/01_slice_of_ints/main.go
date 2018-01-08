package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9}
	myOtherSlice := []int{2, 4, 6, 8}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
