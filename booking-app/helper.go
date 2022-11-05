package main

func ValidateUserInput(firstName string, lastName string, userTickets uint8, remainingTickets uint8) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	//isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidTicketNumber
}
