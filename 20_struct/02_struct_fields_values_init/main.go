package main

import "fmt"

//kreiramo svoj tip
//ima podtip u ovom slucaju to je struct
//tip ima polja (first, secon i age u ovom slucaju) i moze da ima tagove (njih ovde nema)
//deklarisemo promenljive ovog tipa
//promenljive se inicijalizuju ili se postavljaju na zero vrednosti
type person struct {
	first  string
	second string
	age    int
}

func main() {
	p1 := person{"james", "bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.first, p1.second, "godina:", p1.age)
	fmt.Println(p2.first, p2.second, "godina:", p2.age)
}

/*
Go allows us the ability to declare our own types.
 Struct types are declared by composing a fixed set of unique fields together.
 Each field in a struct is declared with a known type.
 This could be a built-in type or another user defined type.

 Once we have a type declared, we can create values from the type
 When we declare variables, the value that the variable represents is always initialized.
 The value can be initialized with a specific value or it can be initialized to its zero value
 For numeric types, the zero value would be 0; for strings it would be empty;
 and for booleans it would be false.

 In the case of a struct, the zero value would apply to all the different fields in the struct.
 Anytime a variable is created and initialized to its zero value, it is idiomatic to use the keyword var.
 Reserve the use of the keyword var as a way to indicate that a variable is being set to its zero value.
 If the variable will be initialized to something other than its zero value,
 then use the short variable declaration operator with a struct literal
*/
