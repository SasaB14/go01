package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Sasa":  "Dobro jutro",
		"Zarko": "Jutric",
	}

	myGreeting["Aca"] = "Morgen"

	fmt.Println(len(myGreeting))
}
