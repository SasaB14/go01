package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//kreiramo map gde je key string i cuva stringove
	words := make(map[string]string)

	//bufio - input&output buffer
	//kreiramo bafer sa NewScanner-om, on za argument uzima Reader
	//a vraca pokazivac na Scanner
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords) //delimo ga po recima

	//skenira sve stvari u skeneru
	for sc.Scan() {
		words[sc.Text()] = "" //daje sve reci kao key a vrednost nista
	}
	if err := sc.Err(); err != nil { //proveravamo gresku, ako greska nije nill stampamo je
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	//idemo kroz sve reci (popravljamo do 200)
	//
	i := 0
	for k, _ := range words { //uzimamo samo key jer vrednost nemamo (pogledaj red 28)
		fmt.Println(k) //stampamo key
		if i == 200 {
			break
		}
		i++
	}
}
