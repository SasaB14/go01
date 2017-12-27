package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers { //radi loop kroz brojeve (ovde 1,2,3,4), kad prodje kroz svaki zove se fukcija kojoj se prosledjuje n
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
