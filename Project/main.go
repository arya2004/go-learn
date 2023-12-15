package main

import (
	"fmt"
	"strings"
)

//package level var. all func have access to them

var conferrenceName string = "GO conf"

const conferrenceTIckets int = 50

var remainingTickets uint = 50

// slice aka List
var bookings []string

func main() {

	// var bookings =  []string{}
	//auto spacing
	//Variables: var name type = value

	fmt.Printf("conferrenceName type %T\n", conferrenceName)

	greetUser()
	//Array
	//var var_name = [size]var_type
	//var var_name = [size]var_type{}
	//var bookings = [50]string{}

	for {
		userName, email, userTickets := getUserInput()

		isValidemail := strings.Contains(email, "@")

		if !isValidemail {
			fmt.Println("not valid email")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Println("Oly this many remaining:", remainingTickets)
			continue
		}

		fmt.Printf("Slice: %v\n", bookings)
		fmt.Printf("First value: %v\n", bookings[0])
		fmt.Printf("Slcie Type: %T\n", bookings)
		fmt.Printf("Slcie LEngth: %v\n", len(bookings))

		bookTickets(userTickets, userName)

		// _ used as blank identifier, unussed vaariables
		// for _, booking := range bookings {
		// 	fmt.Println(booking)
		// }

		printFirstNames()

		if remainingTickets == 0 {
			fmt.Println("All booked!")
			break
		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking app\n", conferrenceName)
	fmt.Println("We have total of", conferrenceTIckets, "and remaining", remainingTickets)
	fmt.Println("Get tickets here")

}

func printFirstNames() {
	for index, booking := range bookings {
		fmt.Println(index, " has", booking)
	}
}

func getUserInput() (string, string, uint) {
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

	return userName, email, userTickets

}

func bookTickets(userTickets uint, userName string) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = userName
	bookings = append(bookings, userName)

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	fmt.Printf("Remaining tickets: %v", remainingTickets)
}
