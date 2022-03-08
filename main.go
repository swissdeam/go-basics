package main

import "fmt"

// pointers

func main() {
	A := 2
	fmt.Println("Value", A) // "A" is still equal 2
	changeValue(&A)         // we must using pointers, 'cause without it, function "changeValue" gets just a copy of "A", not the original "A".
	// in this case, using "&", we say that we give a source of original "A" to the func.
	fmt.Println("Value", A) // after func we see, that fuc used the original "A" and changed it
}

func changeValue(x *int) { // using this pure-func, we must declarate that we use pointer to the original vars
	*x += 2 // in parameters and in body of func

}
