package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"

	//totalTickets
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")

	var bookings [50]string

	var firstname string
	var lastname string
	var userEmail string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstname)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastname)
	fmt.Println("Please enter your email")
	fmt.Scan(&userEmail)
	fmt.Println("Please enter the number of tickets you would like to purchase")
	fmt.Scan(&userTickets)

	// fmt.Println(&userName, &userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstname + " " + lastname

	fmt.Printf("Thank you %s %s, you have successfully booked %d tickets for the %s conference. A confirmation email will be sent to %s\n", firstname, lastname, userTickets, conferenceName, userEmail)

	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available")

}
