package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 30

	fmt.Printf("conferenceName is %T, conferenceTicket is %T, reamaingTickets is %T\n", conferenceName, conferenceTicket, remainingTickets)
	fmt.Printf("welcome to our %v booking application\n", conferenceName)
	fmt.Printf("out of %v tickets %v are available only\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets from here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for their name

	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("enter your ticket number: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets you will receive conformation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf(" %v  remainging tickets for %v \n", remainingTickets, conferenceName)
}
