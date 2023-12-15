package helper

import (
	"fmt"
)

func GreetUser(conferrenceName string, conferrenceTIckets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app\n", conferrenceName)
	fmt.Println("We have total of", conferrenceTIckets, "and remaining", remainingTickets)
	fmt.Println("Get tickets here")

}
