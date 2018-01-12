package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//(p person) je receiver i on ovu funkciju prikljucuje tipu person
//bilo koja vrednost tipa person ima pristup ovoj metodi
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
