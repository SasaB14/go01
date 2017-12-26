package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Whats up Tim or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Whats up Medhi or Marcus")
	case "Julian", "Joe":
		fmt.Println("Hello Joe or Julian")
	default:
		fmt.Println("nemate prijatelje?!?")
	}
}
