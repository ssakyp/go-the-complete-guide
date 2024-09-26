package main

import (
  "fmt"
  "time"
  "errors"
)


func main() {
  userFirstName := getUserData("Please enter your first name: ")
  userLastName := getUserData("Please enter your first name: )
  userBirthDate "= getUserData("Please enter your birthdate (MM/DD/YYYY): ")

  var appUser *user
  appUser, err := newUser(userFirstName, userLastName, birthDate string)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  
  // ... do something awesome with that gathered data

  appUser.outputUserDetails()
  appUser.clearUserData() // copy is cleared
  appUser.outputUserDetails()
                          
}

func getUserData(promptText string) string {
  fmt.Println(promptText)
  var value string
  fmt.Scanln(&value)
  return value
}
