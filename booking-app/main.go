package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTicket int = 50
var conferenceName string = "Go Conference"
var remainingTicket uint = 50
// var bookings = []strings{}  ==> list of slice type string.
// var bookings = make([]map[string]string, 0)  ==> list of slice type map
var bookings = make([]UserData, 0)  

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}
func main() {
	
	fmt.Println("Go by nana")
	

	greetUsers(conferenceName, conferenceTicket, remainingTicket)

	
	firstName, lastName, email, userTickets := getUserInput()

	isValidNAME, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTicket)

	if isValidTickets && isValidNAME && isValidEmail{
		
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		if remainingTicket == 0{
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
		
		firstNames := getFirstName()
		
		fmt.Printf("This are our bookings: %v\n",firstNames)
	}else{
		
		if !isValidNAME{
			fmt.Println("first name or last name you entered is too short.")
		}
		if !isValidEmail{
			fmt.Println("Email you have provided does not contain @ sign.")
		}
		if !isValidTickets{
			fmt.Println("Number of tickets you entered is invalid.")
		}
	}		
	wg.Wait() 
}	


func greetUsers(conferenceName string, conferenceTicket int, remainingTicket uint){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend.")
}

func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint 
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidNAME := len(firstName) >= 2 && len(lastName) >=2 
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidNAME, isValidEmail, isValidTickets
}

func getFirstName() []string {
	firstNames := []string{}		
		for _, booking := range bookings{
			firstNames = append(firstNames, booking.firstName)
		}
		return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTicket = remainingTicket - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v.\n", bookings)	
	
	fmt.Printf("Thank you %v %v for booking %v ticket. you will receive confirmation email at %v.\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n",remainingTicket, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket is for %v %v.\n",userTickets, firstName, lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v.\n",ticket, email)
	fmt.Println("####################################")
	wg.Done()
}