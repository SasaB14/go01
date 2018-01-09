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

	if val, exists := myGreeting[2]; exists {
		fmt.Println("Ta vrednost postoji.")
		fmt.Println("vrednost", val)
		fmt.Println("postoji", exists)
	} else {
		fmt.Println("TA VREDNOST NE POSTOJI")
		fmt.Println("vrednost", val)
		fmt.Println("postoji", exists)
	}

	fmt.Println(myGreeting)
}
