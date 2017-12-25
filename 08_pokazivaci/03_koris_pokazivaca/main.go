package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  //stampa vrednost
	fmt.Println(&a) //stampa adresu

	var b *int = &a

	fmt.Println(b)  //stampa adresu
	fmt.Println(*b) //stampa vrednost

	*b = 11 //vrednost na koju pokazuje pokazivac b promeni na 11
	fmt.Println(a)

	//korisna stvar
	//mozemo proslediti memorijsku adresu umesto nekoliko vrednosti (prosledjujemo referencu)
	//mozemo promeniti vrednost onog sto se cuva na toj adresi

	//ovo cini programe brzim (ne prosledjujemo promenljive vec adrese)

	//u Go-u sve se prosledjuje po vrednosti
	//kada prosledimo memorijsku adresu prosledjujemo vrednost
}
