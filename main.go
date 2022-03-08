package main

import "fmt"

type Person struct { // syntax of creating struct
	name string
	age  int
}

func main() {
	Tom := Person{"G", 34}
	var agePointer *int = &Tom.age // creating pointer for one of struct' field
	var tomPointer *Person = &Tom  // creating pointer for struct as well, through it we can change fields of struct
	*agePointer = 4                // changing field with pointer for that field
	fmt.Println(Tom.age)
	tomPointer.age = 5 // changing field with pointer for ctruct the feild in / syntax | structPointer.{name_field_of_struct} = ...|
	// the same thing as (*tomPointer).age = 5  / idk why yet
	fmt.Println(Tom.age)
}
