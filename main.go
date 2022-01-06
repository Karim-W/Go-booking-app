package main

import (
	// "bookingapp/Classes/resturantCustomer"
	"fmt"
	"strings"
)

func print(arg string) {
	fmt.Println(arg)
}

func queryUser(message string, assignedVar *string) { //would love to use generics but appearently that gonna be there in go 1.18
	fmt.Println(message)
	fmt.Scan(&assignedVar)
}

// func queryUserInt(message string, assignedVar *int) {
// 	fmt.Println(message)
// 	fmt.Scan(&assignedVar)
// }

type ResturantCustomer struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	resturantName := "Memes and giggles"
	const seats = 100
	availableSeats := seats
	fmt.Print("Welcome to ", resturantName, " resturant !!!!\n\n")
	fmt.Println("Total Seats:		", seats)
	fmt.Println("Available Seats:	", availableSeats)
	fmt.Println("Would You like to reseve a table? [Y/N]")
	var response = ""
	fmt.Scan(&response)
	response = strings.ToLower(response)

	if strings.Contains(response, "y") {
		fmt.Println("you said yes")
		//Ask user for their basic info
		var customer ResturantCustomer
		fmt.Println("Please Identify Your First Name")
		fmt.Scanln(&customer.FirstName)
		fmt.Println("Last name please")
		fmt.Scanln(&customer.LastName)
		print("Nice to meet You " + customer.FirstName + " " + customer.LastName)
		fmt.Println("Please Identify Your Age")
		fmt.Scanln(&customer.Age)
		if customer.Age < 18 {
			print("go get your mommy")
		} else {
			print("thanks")
		}
	} else {
		fmt.Println("Thank you, come again")
	}
}
