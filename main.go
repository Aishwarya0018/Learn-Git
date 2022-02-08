package main

import (
	"fmt"
)

func main() {
	//:= only variables
	conferenceName := " Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	//%v placeholder,%T type
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we hAVE total off %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here")
	for {
		var firstname string
		var lastname string
		var email string
		var userTickets uint

		//ask username
		fmt.Println("enter your first name: ")
		fmt.Scan(&firstname)

		fmt.Println("enter your last name: ")
		fmt.Scan(&lastname)

		fmt.Println("enter your email: ")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		//bookings[0] = firstname + " " + lastname
		bookings = append(bookings, firstname+" "+lastname)

		//userTickets = 2
		//	fmt.Println(userName)
		fmt.Printf("Thank you %v for  %v email %v userTickets%v\n", firstname, lastname, email, userTickets)
		fmt.Printf("we have total %v tickets and %v are still available", conferenceName, remainingTickets)

		fmt.Printf("These are all our bookings: %v", bookings)

	}
}
