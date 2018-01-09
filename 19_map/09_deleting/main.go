package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		1: "Dobro jutro",
		2: "dobar dan",
		3: "Dobro Vece",
		4: "Dobro prepodne",
		5: "Dobro popodne",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 4)
	fmt.Println(myGreeting)
}
