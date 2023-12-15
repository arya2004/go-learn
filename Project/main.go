package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferrenceName string = "GO conf"
	const conferrenceTIckets int = 50
	var remainingTickets uint = 50
	//slice aka List
	var bookings []string
	// var bookings =  []string{}
	//auto spacing
	//Variables: var name type = value

	fmt.Printf("conferrenceName type %T\n", conferrenceName)

	fmt.Printf("Welcome to %v booking app\n", conferrenceName)
	fmt.Println("We have total of", conferrenceTIckets, "and remaining", remainingTickets)
	fmt.Println("Get tickets here")

	//Array
	//var var_name = [size]var_type
	//var var_name = [size]var_type{}
	//var bookings = [50]string{}

	for {
		var userName string
		var email string
		var userTickets uint
		//input name
		fmt.Print("\nEnter Usename: ")
		fmt.Scan(&userName)

		fmt.Print("Enter email: ")
		fmt.Scan(&email)

		fmt.Print("Enter TIckes: ")
		fmt.Scan(&userTickets)

		isValidemail := strings.Contains(email, "@")

		if !isValidemail {
			fmt.Println("not valid email")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Println("Oly this many remaining:", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		//bookings[0] = userName
		bookings = append(bookings, userName)

		fmt.Printf("Slice: %v\n", bookings)
		fmt.Printf("First value: %v\n", bookings[0])
		fmt.Printf("Slcie Type: %T\n", bookings)
		fmt.Printf("Slcie LEngth: %v\n", len(bookings))

		for index, booking := range bookings {
			fmt.Println(index, " has", booking)
		}
		// _ used as blank identifier, unussed vaariables
		// for _, booking := range bookings {
		// 	fmt.Println(booking)
		// }

		fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
		fmt.Printf("Remaining tickets: %v", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("All booked!")
			break
		}
	}

}
