package main

import (
  "fmt"
  "time"
)

type user struct {
  firstName string
  lastName string
  birthDate string
  createdAt time.Time
}

func (u *user) outputUserDetails() {
  // ...
   fmt.Println((*u).firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
  u.firtsName = ""
  u.lastName = ""
}

func main() {
  userFirstName := getUserData("Please enter your first name: ")
  userLastName := getUserData("Please enter your first name: )
  userBirthDate "= getUserData("Please enter your birthdate (MM/DD/YYYY): ")

  var appUser user
  appUser = user {
    firstName: userFirstName,
    lastName: userLastName,
    birthDate: userBirthDate,
    createdAt: time.Now(),
  }
  
  // ... do something awesome with that gathered data

  appUser.outputUserDetails()
  appUser.clearUserData() // copy is cleared
  appUser.outputUserDetails()
                          
}

func getUserData(promptText string) string {
  fmt.Println(promptText)
  var value string
  fmt.Scan(&value)
  return value
}
