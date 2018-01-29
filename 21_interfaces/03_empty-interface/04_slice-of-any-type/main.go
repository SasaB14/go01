package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"woof"}, true}
	iva := cat{animal{"meow"}, true}
	shadow := dog{animal{"woof"}, true}
	critters := []interface{}{fido, iva, shadow}
	fmt.Println(critters)
}
