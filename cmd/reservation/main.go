package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//'get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `get` command
	getAllPassengers := getCmd.Bool("all", false, "Get all passenger details")
	getById := getCmd.Int("id", -1, "Passenger ID")
	getPassengerList := getCmd.Bool("users", false, "Get Passenger list")
	//'add' subcommand
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	//'book-ticket' subcommand
	bookTicketCmd := flag.NewFlagSet("book-ticket", flag.ExitOnError)

	//inputs for `book-ticket` command
	bookForId := bookTicketCmd.Int("id", -1, "Enter id of user")

	if len(os.Args) < 2 {
		ExitWithMessage("Expected 'get' (or) 'add' (or) 'book-ticket' subcommands")
	}

	switch os.Args[1] {
	case "get":
		GetAllUserInfo(getCmd, getAllPassengers, getById, getPassengerList)
	case "add":
		AddUser(addCmd)
	case "book-ticket":
		BookTickets(bookTicketCmd, bookForId)
	default:
		fmt.Printf("Invalid command - %s. Expected commands 'get' (or) 'add' (or) book-ticket ", os.Args[1])
	}

}
