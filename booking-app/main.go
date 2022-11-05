package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets uint8 = 50

var conferenceName = "Go Conference"
var remainingTickets uint8 = 50
var bookings = make([]map[string]string, 0)

//var bookings = [50]string{"Mordi", "Knix"}
//var bookings [50]string  <-- Array \/ slice
// bookings := []string{}

// VIDEO
// https://youtu.be/yyUHQIec83I?t=8926

// Where does the compiler start, the entrypoint
func main() {

	greetUsers()

	for { //remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, userTickets := getUserInput()
		isValidName, isValidTicketNumber := ValidateUserInput(firstName, lastName, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the outer for-loop
				fmt.Println("Our conference is sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname or lastname is too short.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets are invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { //the index, _, we are not using it
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, uint8) {
	var firstName string
	var lastName string
	//		var email string
	var userTickets uint8

	// ask user for name
	fmt.Println("Enter you first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name:")
	fmt.Scan(&lastName)

	//		fmt.Println("Enter you email:")
	//		fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	// create a map to store bookings
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}
