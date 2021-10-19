# mock-reservations-cli

A golang application to mock the train reservation system 

**Features**

1. View all passenger details
2. Add user to the train passengers list
3. Book tickets for a passenger
   
**Usage**

1. Move to cmd/reservation directory
2. Run `go mod init` & `go mod tidy` to create `go.mod` & `go.sum` files
3. Run the cmd.exe file with flags
  - get 
    - all (to view all passengers details)
    - id (to view user by id)
    - users (to view passengers list alone)
  - add (to add new user)
  - book-ticket 
    - id (to book ticket for a user)
  
