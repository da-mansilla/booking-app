package main

import "fmt"

func main() {
	const conferenceTickets = 50
	conferenceName := "Go Conference"
	remainingTickets := 50

	fmt.Println("Welcome to",conferenceName,"booking application")
	fmt.Println("We have total of ",conferenceTickets, "tickets and", remainingTickets,"are stil available")
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var numberTickets int 
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Enter the numbers of tickets that you want to: ")
	fmt.Scan(&numberTickets)


	fmt.Printf("%v booked %v tickets\n",userName,numberTickets)


}
