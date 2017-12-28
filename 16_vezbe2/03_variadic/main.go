package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57, 100} // ovo je slice
	n := average(data...)                          //rasuturamo ga na delove sa data...
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	max := 0.0
	for _, v := range sf { //_ je index, v je vrednost
		if max < v {
			max = v
		}

	}
	return max
}
