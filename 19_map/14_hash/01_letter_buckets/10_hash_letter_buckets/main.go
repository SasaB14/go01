package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// daj knjigu moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	//skeniraj stranu
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close() //defer - pre nego sto se ova funkcija zavrsi zatvori Body

	//podesi split funkciju za operaciju skeniranja
	scanner.Split(bufio.ScanWords) //delimo ga po recima

	//kreiramo slice da drzi broj
	buckets := make([]int, 200)

	//prodji kroz sve reci
	for scanner.Scan() {
		n := hashBucket(scanner.Text()) //funkcija je nize
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	fmt.Println("***************")
	for i := 28; i <= 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	}

}

//uzima rec i prvo slovo pretvara u int
func hashBucket(word string) int {
	return int(word[0])
}
