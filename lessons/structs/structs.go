package main

import (
  "fmt"
  "github.com/ssakyp/structs/user"
)


func main() {
  userFirstName := getUserData("Please enter your first name: ")
  userLastName := getUserData("Please enter your first name: )
  userBirthDate "= getUserData("Please enter your birthdate (MM/DD/YYYY): ")

  var appUser *user.User
  appUser, err := user.New(userFirstName, userLastName, birthDate string)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  
  // ... do something awesome with that gathered data

  appUser.OutputUserDetails()
  appUser.ClearUserData() // copy is cleared
  appUser.OutputUserDetails()
                          
}

func getUserData(promptText string) string {
  fmt.Println(promptText)
  var value string
  fmt.Scanln(&value)
  return value
}
