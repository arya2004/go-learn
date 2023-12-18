package main

import (
	"fmt"
)

//Go doesnt have \0 at last of string.
//variable are descriptors containing length values

func main() {
	var s string = "elite"
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", []rune(s))
	fmt.Printf("%v\n", []rune(s))

}
