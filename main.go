package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	fmt.Println(conferenceName)
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string
	var userTickets int
	// ask user for their name

	//fmt.Scan(&userName)
	fmt.Println(remainingTickets)

	fmt.Println(&remainingTickets)
	
	
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)


}

