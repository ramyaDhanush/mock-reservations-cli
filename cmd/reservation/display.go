package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"strings"
)

// ShowPassengerTable - Display a table of user data
func ShowPassengerTable() {
	tbl := table.New("ID", "Name", "Age", "Gender", "Aadhaar No.",  "Number of Travels")
	tbl.WithHeaderFormatter(func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	})

	passengers := GetPassengerDetails()

	for _, passenger := range passengers {
		tbl.AddRow(passenger.Id, passenger.Name, passenger.Age, passenger.Gender, passenger.AadhaarNo, len(passenger.TravelHistory))
	}

	tbl.Print()
}

// PrintPassengerDetails - Display all user details in Table form
func PrintPassengerDetails(passenger Passenger) {
	fmt.Println()
	fmt.Print("PASSENGER NAME - ")
	color.Set(color.FgMagenta)
	fmt.Printf("%s\n", strings.ToUpper(passenger.Name))
	color.Unset()
	fmt.Printf("ID : %d, AGE : %d, GENDER : %s, AADHAAR NO. : %s\n", passenger.Id,  passenger.Age, passenger.Gender, passenger.AadhaarNo)
	//fmt.Println()
	PrintPassengerTravelHistory(passenger.TravelHistory)
}

// PrintPassengerTravelHistory - Display travel history(s) in Table form
func PrintPassengerTravelHistory(history []TravelHistory) {
	if len(history) < 1 {
		color.Set(color.FgCyan)
		fmt.Printf("NO TRAVEL RECORDS\n")
		color.Unset()
		return
	}
	color.Set(color.FgCyan)
	fmt.Printf("TRAVEL RECORDS\n")
	color.Unset()
	tbl := table.New( "Train name", "Train number", "Date", "Time", "From", "To", "Travel Class").WithPadding(4)
	for _, travel := range history {
		tbl.AddRow(travel.TrainName, travel.TrainNumber, travel.Date,travel.Time, travel.Location.From, travel.Location.To, travel.TravelClass)
	}

	tbl.Print()
}

// ExitWithMessage - Display a message & terminate the program
func ExitWithMessage(msg string) {
	color.Set(color.FgRed)
	fmt.Printf(msg)
	color.Unset()
	os.Exit(1)
}

// DisplayMessage - Display a message
func DisplayMessage(msg string) {
	color.Set(color.FgRed)
	fmt.Println(msg)
	color.Unset()
}