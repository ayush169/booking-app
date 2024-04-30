package main

import "fmt"

func main() {
	var conferenceName = "GopherCon"

	const totalTickets = 50
	var remainingTickets = 30

	//print types of variables
	fmt.Printf("conferenceName: %T totalTickets: %T remainingTickets: %T\n", conferenceName, totalTickets, remainingTickets)

	// fmt.Println("Welcome to" + conferenceName + "2021!")
	// fmt.Println("Welcome to", conferenceName, "2021!")
	fmt.Printf("Welcome to %v 2021!\n", conferenceName)
	fmt.Println("Total tickets available:", totalTickets, "Remaining tickets:", remainingTickets)
	// fmt.Printf("Do you want to buy a ticket? (y/n): %d\n", totalTickets)
	// fmt.Println("yes")

	var username string

	username = "John Doe"
	fmt.Println("Welcome", username)

}
