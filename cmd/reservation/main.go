package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func HandleGet(getCmd *flag.FlagSet, all *bool, name *string) {
	// to parse cmdline arguments after subcommand get
	err := getCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if !*all && *name == "" { //invalid input
		fmt.Print("Required : --all or --name for get command")
		getCmd.PrintDefaults()
		os.Exit(1) // Exit execution
	}

	if *all {
		passengers := GetPassengerDetails()

		for _, passenger := range passengers {
			PrintPassengerDetails(passenger)
			PrintPassengerTravelHistory(passenger)
		}
		return
	}

	if *name != "" {
		passengers := GetPassengerDetails()
		name := *name

		isFound := false


		for _, passenger := range passengers {
			if name == passenger.Name {
				isFound = true
				PrintPassengerDetails(passenger)
				PrintPassengerTravelHistory(passenger)
			}
		}
		if !isFound {
			fmt.Printf("**** NOT FOUND : PASSENGER \"%s\" ****\n", strings.ToUpper(name))
		}
	}
}

func ValidateAddCommand(addCmd *flag.FlagSet, name *string, age *int, gender *string, aadhaar *string) {
	err := addCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if *name == "" || *age == 0 || *gender == "" || *aadhaar == "" {
		fmt.Println("All fields [name, from, to] are required for adding user")
		addCmd.PrintDefaults()
		os.Exit(1)
	} else {
		fmt.Println("Passenger details added successfully.....")
	}
}

// HandleAddUser - To add passenger detail
func  HandleAddUser(addCmd *flag.FlagSet, name *string, age *int, gender *string, aadhaar *string )  {
	ValidateAddCommand(addCmd, name, age, gender, aadhaar)

	newPassenger := Passenger{
		Name : *name,
		Age : *age,
		Gender : *gender,
		AadhaarNo: *aadhaar,
	}

	passengers := GetPassengerDetails()
	passengers = append(passengers, newPassenger)

	WritePassengerDetails(passengers)
}

// HandleTravelHistory - To add travel details to a passenger
func HandleTravelHistory(name, trainName, trainNumber, travelDate, travelDay, travelTime, from, to, class *string) {

	newLocation := Location{
		From: *from,
		To : *to,
	}
	newTravelHistory := TravelHistory{
		TrainName: *trainName,
		TrainNumber: *trainNumber,
		Date: *travelDate,
		Day: *travelDay,
		Time: *travelTime,
		Location: newLocation,
		TravelClass: *class,
	}

	passengers := GetPassengerDetails()

	for i, passenger := range passengers {
		if passenger.Name == *name {
			passenger.TravelHistory = append(passenger.TravelHistory, newTravelHistory)
			passengers[i] = passenger
			break
		}
	}

	WritePassengerDetails(passengers)
}

func main() {

	//'get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `get` command
	getAllPassengers := getCmd.Bool("all", false, "Get all passenger details")
	getByName := getCmd.String("name", "", "Passenger name")

	//'add' subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// inputs for 'add' command
	addName := addCmd.String("name", "", "Passenger Name")
	addAge := addCmd.Int("age", 0, "Passenger Age")
	addGender := addCmd.String("gender", "", "Passenger Gender")
	addAadhaar := addCmd.String("aadhaar", "", "Passenger Aadhaar Number")
	addTrainName := addCmd.String("train-name","", "Train Name")
	addTrainNumber := addCmd.String("train-no","", "Train Number")
	addTravelDate := addCmd.String("date","", "Date of Travel")
	addTravelDay := addCmd.String("Day","", "Day of Travel")
	addTravelTime := addCmd.String("Time","", "Time of Travel")
	addFrom := addCmd.String("from","", "Travel from")
	addTo := addCmd.String("to","", "Travel to")
	addTravelClass := addCmd.String("class","", "Class of Travel")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAllPassengers, getByName)
	case "add":
		HandleAddUser(addCmd, addName, addAge, addGender, addAadhaar)
		HandleTravelHistory(addName,  addTrainName, addTrainNumber, addTravelDate, addTravelDay, addTravelTime, addFrom, addTo, addTravelClass)
	default:
		fmt.Printf("Invalid command - %s. Expected commands 'get' (or) 'add' ", os.Args[1])
	}

}
