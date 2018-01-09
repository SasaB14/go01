package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Sasa":  "Dobro jutro",
		"Zarko": "Jutric",
	}

	myGreeting["Aca"] = "Morgen"
	fmt.Println(myGreeting)
	myGreeting["Aca"] = "Good mornin"
	fmt.Println(myGreeting)
}
