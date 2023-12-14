package main

import (
	"fmt"
)

func main() {
	var conferrenceName string = "GO conf"
	const conferrenceTIckets int = 50
	var remainingTickets int = 50
	//auto spacing
	//Variables: var name type = value

	fmt.Printf("conferrenceName type %T\n", conferrenceName)

	fmt.Printf("Welcome to %v booking app\n", conferrenceName)
	fmt.Println("We have total of", conferrenceTIckets, "and remaining", remainingTickets)
	fmt.Println("Get tickets here")

	var userName string
	var userTickets int
	//input name

	fmt.Scan(&userName)

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
