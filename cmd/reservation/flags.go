package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
)

// GetAllUserInfo - To get passenger info based on flags
func GetAllUserInfo(getCmd *flag.FlagSet, all *bool, id *int, users *bool) {
	// to parse cmdline arguments after subcommand get
	err := getCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if !*all && *id == -1 && !*users { //invalid input
		getCmd.PrintDefaults()
		ExitWithMessage("Required : --all (or) --name (or) --users for get command")
	}

	if *all {
		passengers := GetPassengerDetails()
		ShowPassengerTable()

		for _, passenger := range passengers {
			PrintPassengerDetails(passenger)
		}
		return
	}

	if *users {
		ShowPassengerTable()
		return
	}

	if *id != -1 {
		passengers := GetPassengerDetails()
		userId := *id

		isFound := false

		for _, passenger := range passengers {
			if userId == passenger.Id {
				isFound = true
				PrintPassengerDetails(passenger)
			}
		}
		if !isFound {
			fmt.Printf("**** NOT FOUND : PASSENGER \"%d\" ****\n", userId)
		}
	}
}

// AddUser - To register a user
func AddUser(addCmd *flag.FlagSet) {
	newPassenger, err := GetUserDetailsFromUser()
	if err != nil {
		var proceed string
		fmt.Println("Please enter valid passenger data")
		fmt.Print("Do you want to try again? [Y/N] ")
		fmt.Scanln(&proceed)
		proceed = strings.ToLower(proceed)
		if proceed == "y" {
			newPassenger, err = GetUserDetailsFromUser()
			if err != nil {
				ExitWithMessage("Please try again with correct info! Thanks")
			}
		} else {
			ExitWithMessage("Thanks for using our service!! ")
		}
	}
	passengers := GetPassengerDetails()
	passengers = append(passengers, newPassenger)

	WritePassengerDetails(passengers)

	color.Set(color.FgGreen)
	fmt.Printf("Successfully Added Info of Passenger %s [ID : %d]\n", newPassenger.Name, newPassenger.Id)
	color.Unset()

}

// BookTickets - To book a ticket under a user id
func BookTickets(bookTicketCmd *flag.FlagSet,  id *int) {
	ValidateBookTicketCommand(bookTicketCmd, id)

	fmt.Println("----------------------------------------------------")
	fmt.Println("--Welcome to Ticket Booking. Provide below details--")
	fmt.Println("----------------------------------------------------")
	fmt.Println()

	var trainName, date, time, fromPlace, toPlace, classOfTravel string

	trainNameReader := bufio.NewScanner(os.Stdin) // to get a line of input
	fmt.Print("Train Name [cherran express / chennai express] : ")
	trainNameReader.Scan()
	trainName += trainNameReader.Text()


	trainName = strings.ToLower(trainName)
	trainNumber, ok := GetTrainNumber(trainName)
	if ok{
		fmt.Println("Your train number                : ", trainNumber)
	} else {
		ExitWithMessage(fmt.Sprintf("Sorry no train with name %s exists. Try again with correct info", trainName))
	}
	fmt.Print("Enter date of travel [dd-mm-yyyy]      : ")
	fmt.Scanln(&date)
	fmt.Println("Available Timings                    : ", GetAvailableTimeSlots(trainName))
	fmt.Print("Enter preferred time                   : ")
	fmt.Scanln(&time)
	fmt.Print("Boarding Place                         : ")
	fmt.Scanln(&fromPlace)
	fmt.Print("Destination                            : ")
	fmt.Scanln(&toPlace)
	fmt.Println("Available Classes [AC, FC, SC]       : ")
	fmt.Print("Class of Travel                        : ")
	fmt.Scanln(&classOfTravel)

	if ValidateBookingInfo(trainName, date, time, fromPlace, toPlace, classOfTravel){
		fmt.Println(trainName, date, time, fromPlace, toPlace, classOfTravel)
		newTravelHistory := TravelHistory{
			TrainName   : trainName,
			TrainNumber : trainNumber,
			Date        : date,
			Time        : time,
			Location    : Location{
				From : fromPlace,
				To : toPlace,
			},
			TravelClass : classOfTravel,
		}

		passengers := GetPassengerDetails()
		userId := *id
		newPassengerList := []Passenger{}
		currentUser := Passenger{}
		for _, passenger := range passengers {
			if userId == passenger.Id {
				passenger.TravelHistory = append(passenger.TravelHistory, newTravelHistory)
				currentUser = passenger
			}
			newPassengerList = append(newPassengerList, passenger)
		}

		fmt.Println(currentUser.TravelHistory)

		color.Set(color.FgGreen)
		fmt.Printf("Successful!! Booking for passenger %s (ID : %d) from %s to %s on %s at %s under class %s in %s is CONFIRMED",
			currentUser.Name, currentUser.Id, fromPlace, toPlace, date, time, classOfTravel, trainName)
		color.Unset()
		WritePassengerDetails(newPassengerList)

	} else {
		ExitWithMessage("Please enter valid info for booking.. Try again")
	}
}


//Utility for flag

// GetUserDetailsFromUser - To prompt & receive user info
func GetUserDetailsFromUser() (Passenger, error) {
	var name, gender, aadhaarNo string
	var id, age int


	fmt.Println("---------------------------------------------------")
	fmt.Println("--Provide all mandatory (*) info to add passenger--")
	fmt.Println("---------------------------------------------------")
	fmt.Println("")
	fmt.Print("ID*  (Unique number)           : ")
	fmt.Scanln(&id)
	fmt.Print("NAME*                          : ")
	fmt.Scanln(&name)
	fmt.Print("AGE*                           : ")
	fmt.Scanln(&age)
	fmt.Print("GENDER* [M/F/O]                : ")
	fmt.Scanln(&gender)
	fmt.Print("AADHAAR NO* (xxxx-xxxx-xxxx)   : ")
	fmt.Scanln(&aadhaarNo)

	if ValidateUserDetails(id, name, age, gender, aadhaarNo){
		newPassenger := Passenger{
			Id : id,
			Name:      name,
			Age:       age,
			Gender:    gender,
			AadhaarNo: aadhaarNo,
		}
		return newPassenger, nil
	}
	return Passenger{}, errors.New("Invalid data")
}

