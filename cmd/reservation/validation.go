package main

import (
	"flag"
	"os"
	"regexp"
	"strings"
	"strconv"
)

// ValidateBookTicketCommand - Perform sanity check to validate book-ticket command
func ValidateBookTicketCommand(bookTicketCmd *flag.FlagSet, id *int) {
	err := bookTicketCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}
	passengers := GetPassengerDetails()
	userId := *id

	isFound := false
	for _, passenger := range passengers {
		if userId == passenger.Id {
			isFound = true
		}
	}

	if *id == -1 && isFound {
		bookTicketCmd.PrintDefaults()
		ExitWithMessage("Valid 'id' is required to book ticket")
	}
}

// ValidateUserDetails - Perform sanity check for validating user/passenger info
func ValidateUserDetails(id int, name string, age int, gender, aadhaarNo string) bool {
	gender = strings.ToLower(gender)
	passengers := GetPassengerDetails()
	for _, passenger := range passengers {
		if id == passenger.Id {
			DisplayMessage("User ID already exists")
			return false
		}
	}

	return id != 0 && name != "" &&
		(age != 0 && age > 0) &&
		(gender == "f" || gender == "m" || gender == "o") &&
		IsValidAadhaarNumber(aadhaarNo)
}

// IsValidAadhaarNumber - Validate aadhaar number
func IsValidAadhaarNumber(str string) bool {
	strings.Trim(str, " ")
	_, ok := regexp.MatchString("[0-9]{4}-[0-9]{4}-[0-9]{4}", str)

	passengers := GetPassengerDetails()
	for _, passenger := range passengers {
		if str == passenger.AadhaarNo {
			DisplayMessage("Aadhaar Number already exists")
			return false
		}
	}
	return len(str) == 14 && ok == nil
}

// GetStringToIntSlice - Convert string slice to int slice
func GetStringToIntSlice(stringSlice []string) []int {
	var intSlice []int = []int{}

	for _, i := range stringSlice {
		value,  err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, value)
	}

	return intSlice
}

// IsValidDate - Validate Date of booking
func IsValidDate(date string) bool{
	if date == "" {
		return false
	}
	dateArr := strings.Split(date, "-")
	if len(dateArr) < 3{
		return false
	}

	dateArrDec := GetStringToIntSlice(dateArr)
	isYearValid := dateArrDec[2] >= 2021
	isMonthValid := dateArrDec[1] > 0 && dateArrDec[1] <= 12

	var maxDays int
	switch dateArrDec[1] {
	case 1, 3, 5, 7, 8, 10, 12:
		maxDays = 31
	case 4, 6, 9, 11:
		maxDays = 30
	case 2:
		if dateArrDec[2]%4 == 0 && dateArrDec[2]%100 == 0 && dateArrDec[2]%400 !=0 { //leap year
			maxDays = 29
		} else {
			maxDays = 28
		}
	}

	isDateValid := dateArrDec[0] > 0 && dateArrDec[0] <= maxDays
	return isDateValid && isMonthValid && isYearValid
}

// IsValidTimeSlot - Validate Time Slot for booking
func IsValidTimeSlot(timeSlot , trainName string) bool {
	for _, slot := range AvailableSlots[trainName] {
		if slot == timeSlot {
			return true
		}
	}
	return false
}

// ValidateBookingInfo - To Validate all details required for booking a ticket
func ValidateBookingInfo(trainName, date, time, fromPlace, toPlace, classOfTravel string) bool{
	trainName = strings.ToLower(trainName)
	classOfTravel = strings.ToLower(classOfTravel)

	return (trainName == "cherran express" || trainName == "chennai express") &&
		IsValidDate(date) &&
		IsValidTimeSlot(time, trainName) &&
		fromPlace != "" && toPlace != "" &&
		(classOfTravel == "ac" || classOfTravel == "fc" || classOfTravel == "sc")
}