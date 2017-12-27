package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // ovo je slice
	n := average(data...)                     //rasuturamo ga na delove sa data...
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf { //_ je index, v je vrednost
		total += v
	}
	return total / float64(len(sf))
}
