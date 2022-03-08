package main

import "fmt"

type library []string // declarating type library
//type department []string

func (l library) printAll() { // function just for library type, it will work only with library vars
	for _, val := range l { // Printing all values in massive which came in function
		fmt.Println(val)
	}
}

func main() {

	var cityLib library = library{"book1", "book2", "book3"} // creating library var with 3 values in it
	//var catalog department = department{"matter", "matter2", "matter3"}
	cityLib.printAll() // use func wich was made only for library vars
	//catalog.printAll() doesnt work with printAll
}
