package main

import (
	"fmt"
	"strings"
)

// VIDEO
// https://youtu.be/yyUHQIec83I?t=5974

// Where does the compiler start, the entrypoint
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	fmt.Printf("Welcome to our %v booking system\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{"Mordi", "Knix"}
	//var bookings [50]string  <-- Array \/ slice
	var bookings []string
	// bookings := []string{}

	for { //remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		//		var email string
		var userTickets uint

		// ask user for name
		fmt.Println("Enter you first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter you last name:")
		fmt.Scan(&lastName)

		//		fmt.Println("Enter you email:")
		//		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want:")
		fmt.Scan(&userTickets)

		if userTickets <= uint(remainingTickets) {
			remainingTickets = remainingTickets - uint8(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings { //the index, _, we are not using it
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all out bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the outer for-loop
				fmt.Println("Our conference is sold out")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
		}
	}

}
