package main

import "fmt"

func main() {
  firstName := getUserData("Please enter your first name: ")
  lastName := getUserData("Please enter your first name: )
  birthDate "= getUserData("Please enter your birthdate (MM/DD/YYYY): ")

  // ... do something awesome with that gathered data

  outputUserDetails(firstName, lastName, birthDate)
                          
}

func outputUserDetails(firstName, lastName, birthDate string) {
  // ...
   fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
  fmt.Println(promptText)
  var value string
  fmt.Scan(&value)
  return value
}
