package main


import (
	"fmt"
	"github.com/fatih/color"
	"strings"
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