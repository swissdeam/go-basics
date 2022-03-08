package main

import "fmt"

// pointers

func main() {
	x := "value of x"
	a := &x         // & means where "a" take the data; in this case "a" takes data from "x"
	fmt.Println(a)  // without "*", "a" equals addres of "x"
	fmt.Println(*a) // with "*", "a" means value of source(in this case source is "x")

}
