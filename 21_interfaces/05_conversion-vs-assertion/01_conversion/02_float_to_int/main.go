package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(int(y) + x) // konverzija: float64 u int
}