package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)        //funkcija dobija kopiju vrednosti na pravu vrednost za X
	fmt.Println(x) //x je i dalje 5
}
