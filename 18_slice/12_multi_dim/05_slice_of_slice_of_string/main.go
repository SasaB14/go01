package main

import "fmt"

func main() {
	var records [][]string

	//djak broj 1
	student1 := make([]string, 4)
	student1[0] = "Sasa"
	student1[1] = "Brankovic"
	student1[2] = "100.00"
	student1[3] = "74.00"

	//sacuvaj podatke
	records = append(records, student1)

	//djak broj 2
	student2 := make([]string, 4)
	student2[0] = "Zarko"
	student2[1] = "Mavrenovic"
	student2[2] = "110.00"
	student2[3] = "96.00"

	//sacuvaj podatke
	records = append(records, student2)

	fmt.Printf("%T\n", records)
	fmt.Println(records)
}
