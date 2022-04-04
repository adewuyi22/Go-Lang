package main
import (  "fmt"
			"String"

)

func bookingApp() {

	 conferenceName := "GO CONFERENCE"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	//using slice method to declare an array
	 bookings := [] string {} 
  
	var firstName string 
	var lastName string
	var email string  
	var userTicket uint
	

	fmt.Printf("\n--------------------------------------------------------- \nWelcome to %v Booking Application \n",conferenceName)
	fmt.Printf("We have a total of %v ticket(s) available for booking.\n",  conferenceTickets)
	fmt.Println("Input your details below here to attend. \n---------------------------------------------------------")

	i := 0
	Here: // label ends with ":"
	//to be continue from the last number of bookings 
	for {

		//asking for user input
	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("\nEnter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("\nEnter your email: ")
	fmt.Scan(&email)

	fmt.Println("\nEnter the number of tickets you want to book: ")
	fmt.Scan(&userTicket)	


	//Validating UserInput
	isValidName := len(firstName) => 2 && len(lastName) => 2 
	isValidName := strings.Contains(email, "@")  
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	

	if (userTicket > remainingTickets) {
		fmt.Printf("\nSorry, Only %v ticket(s) is available for booking, Kindly Rebook. \n", remainingTickets)

		//fmt.Println(i)
		i++
		goto Here // jump back to label "Here"

	}  else {
		remainingTickets = remainingTickets - userTicket
			//bookings [0] = firstName + "" + lastName
	bookings = append(bookings, firstName + " " + lastName + ",")

	/*
		fmt.Printf("The whole slice : %v  \n", bookings)
		fmt.Printf("The first value : %v  \n", bookings[0])
		fmt.Printf("The slice length : %v  \n", len(bookings))
	*/
		fmt.Printf("\nWelcome %v %v you have successfully booked %v ticket(s).", firstName, lastName, userTicket)
		fmt.Printf("\nYou will receive a confirmation mail at %v \nThanks for booking with us \n", email)
		fmt.Printf("%v tickets is remaining for booking at %v \n", remainingTickets, conferenceName)
		fmt.Printf("These are all our bookings: %v \n", bookings)
	
	} 

}


	}

func main(){
	bookingApp()
}
