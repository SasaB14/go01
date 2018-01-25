package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(n)
	// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
	//type IntSlice []int
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

}
