package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	frendly bool
}

type cat struct {
	animal
	frendly bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	iva := dog{animal{"cat"}, true}
	specs(fido)
	specs(iva)
}
