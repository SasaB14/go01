package main

import "fmt"

func main() {
	sumaTriPet := 0
	sumaOst := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			sumaTriPet = sumaTriPet + i
		} else if i%5 == 0 {
			sumaTriPet = sumaTriPet + i
		} else {
			sumaOst = sumaOst + i
		}
	}
	fmt.Println("Suma brojeva deljivih sa 3 i 5 od 1 do 1000 je: ", sumaTriPet)
	fmt.Println("Suma brojeva koji nisu deljivi sa 3 i 5 od 1 do 1000 je: ", sumaOst)
}
