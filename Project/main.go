package main

import (
	"fmt"
	"os"
)

func Args() {
	if len(os.Args) > 1 {
		fmt.Println("HEllo, ", os.Args[1])
		return
	} else {
		fmt.Println("There is no args")
		return
	}

}

func main() {
	Args()
}
