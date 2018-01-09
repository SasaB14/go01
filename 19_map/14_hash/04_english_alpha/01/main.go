package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http.Get daje response i error, ako imamo error prekidamo program
	// response je struct i jedno od polja je Body tipa je io.ReadCloser
	//ReadCloser je interfejs koji implementira Reader i Closer
	//Reader je jos jedan interfejs
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//ReadAll uzima za argument Reader i iscitava sve
	bs, _ := ioutil.ReadAll(res.Body) //read all vraca slice of bytes
	str := string(bs)                 //konvertujemo ih stringove
	fmt.Println(str)
}
