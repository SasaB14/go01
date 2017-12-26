package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Whats up Daniel")
	case "Medhi":
		fmt.Println("Whats up Medhi")
	case "Jenny":
		fmt.Println("Hello Jenny")
	default:
		fmt.Println("nemate prijatelje?!?")
	}
}

/*
nema podrazumevanog fallthrough-a, on je opcion
-- moze se staviti podrazumevani tako sto se eksplicitno postavi
-- break nije potreban kao u C++-u
*/
