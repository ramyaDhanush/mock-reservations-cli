package main

import (
	"strings"
)

var AvailableSlots map[string][]string = map[string][]string {
	"cherran express" : []string{"07:00:10", "09:21:00", "13:30:00", "16:90:10"},
	"chennai express" : []string{"19:12:00", "24:10:00", "02:23:90", "06:45:00"},
}

var TrainNumberLookup map[string]string = map[string]string {
	"cherran express" : "204566",
	"chennai express" : "221202",
}

// GetTrainNumber - Return a valid train number if exists
func GetTrainNumber(name string) (string, bool) {
	train, ok :=  TrainNumberLookup[name]

	var result string

	if ok {
		result = train
	} else {
		result = "Sorry, No Train Found with that name"
	}

	return result, ok
}

// GetAvailableTimeSlots - Return a time slot if available
func GetAvailableTimeSlots(name string) string{
	slots, ok :=  AvailableSlots[name]

	var result string

	if ok {
		result = strings.Join(slots, " ")
	} else {
		result = "Sorry, No Time Slots!"
	}

	return result
}

