package user

import (
  "fmt"
  "errors"
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
