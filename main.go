package main

import "fmt"

// pointers

func main() {
	x := new(int)            // Creating pointer for nameless var with int-type
	fmt.Println("value", *x) //value 0
	*x = 6                   // changing value('cause there is "*") of nameless var
	fmt.Println("value", *x) // value 8

}
