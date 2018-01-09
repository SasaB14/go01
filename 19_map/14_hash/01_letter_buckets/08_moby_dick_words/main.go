package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//NewScanner uzima Reader - Newreader uzima string i vraca reader
	//vraca pokazivac na scanner
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close() //defer - pre nego sto se ova funkcija zavrsi zatvori Body

	// podesavamo split funciju da bi mogli da skeniramo reci.
	// ScanWords je split funkcija za Scanner koja vraca
	// svaku rec odvojenu space-om i brise sve space-ove oko nje.
	scanner.Split(bufio.ScanWords)

	// Prebroj reci
	//Scan() vraca false kada se skeniranje zavrsi
	//Err metoda vraca bilo koju gresku koja se pojavila kod skeniranja
	// osim ako je u pitanju io.EOF, tad Err vraca nil
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
