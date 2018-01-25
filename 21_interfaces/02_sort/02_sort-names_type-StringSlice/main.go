package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	//StringSlice
	//StringSlice attaches the methods of Interface to []string, sorting in increasing order.
	//definicija - type StringSlice []string

	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

// https://golang.org/pkg/sort/#Sort
