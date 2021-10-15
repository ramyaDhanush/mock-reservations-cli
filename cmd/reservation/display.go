package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"github.com/rodaine/table"
)


func PrintPassengerDetails(passenger Passenger) {
	color.Set(color.FgRed)
	fmt.Printf("PASSENGER NAME - %s\n", strings.ToUpper(passenger.Name))
	color.Unset()
	fmt.Printf("PASSENGER INFO - AGE : %d, GENDER : %s, AADHAAR NO. : %s\n", passenger.Age, passenger.Gender, passenger.AadhaarNo)
}

func PrintPassengerTravelHistory(passenger Passenger) {
	fmt.Println("_____________________________________")
	for i, history := range passenger.TravelHistory {
		color.Cyan(fmt.Sprintf("Travel Entry : %d", i+1))
		//color.Set(color.FgYellow)
		fmt.Printf("TRAIN NAME : %s, TRAIN NUMBER : %s\n", history.TrainName, history.TrainNumber)
		//color.Unset()
		fmt.Printf("DATE : %s, DAY : %s, TIME : %s\n", history.Date, history.Date, history.Time)
		fmt.Printf("LOCATION  : Travelled from %s to %s\n", history.Location.From, history.Location.To)
		fmt.Printf("CLASS of TRAVEL : %s\n", history.TravelClass)
	}
	fmt.Println("_____________________________________")
}

func FormatTravelHistory(history []TravelHistory) table.Table {
	tbl := table.New("Name", "Number", "Date", "Day", "Time", "From", "To", "Travel Class")

	for i, travel := range history {
		tbl.AddRow(i, travel.TrainName, travel.TrainNumber, travel.Date, travel.Day, travel.Time, travel.Location, travel.TravelClass)
	}

	return tbl
}

func ShowPassengerTable() {
	tbl := table.New("ID", "Name", "Age", "Gender", "Aadhaar No.",  "Number of Travels")

	passengers := GetPassengerDetails()

	for _, passenger := range passengers {
		tbl.AddRow(passenger.Id, passenger.Name, passenger.Age, passenger.Gender, passenger.AadhaarNo, len(passenger.TravelHistory))
	}

	tbl.Print()
}