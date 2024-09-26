package main

import (
  "fmt"
  "time"
  "errors"
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

// constructor pattern
func newUser(firstName, lastName, birthDate string) (*user, error) {
  if firtsName == "" || lastName == "" || birthDate == "" {
      return nil, errors.New("First name, last name, birth date is required!")
  }
  return &user {
    firstName: firstName,
    lastName: lastName,
    birthDate: birthDate,
    createdAt: time.Now(),
  }, nil
}
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
