package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Location struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type TravelHistory struct {
	TrainName   string   `json:"train-name,omitempty"`
	TrainNumber string   `json:"train-number,omitempty"`
	Date        string   `json:"date,omitempty"`
	Time        string   `json:"time,omitempty"`
	Location    Location `json:"location,omitempty"`
	TravelClass string   `json:"travel-class,omitempty"`
}

type Passenger struct {
	Id            int          `json:"id"`
	Name          string          `json:"name"`
	Age           int             `json:"age"`
	Gender        string          `json:"gender"`
	AadhaarNo     string          `json:"aadhaar-no"`
	TravelHistory []TravelHistory `json:"travel-history,omitempty"`
}

// GetPassengerDetails - To get the list of existing passenger info
func GetPassengerDetails() (passenger []Passenger) {
	fileBytes, err := ioutil.ReadFile("../api/passenger.json")
	if err != nil {
		fmt.Println(err, fileBytes)
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &passenger)

	if err != nil {
		panic(err)
	}
	return passenger
}

// WritePassengerDetails - To Write passenger details to json file
func WritePassengerDetails(passengers []Passenger) {
	passengerBytes, err := json.Marshal(passengers)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("../api/passenger.json", passengerBytes, 0644)
	if err != nil {
		panic(err)
	}
}
