package user

import (
  "fmt"
  "errors"
  "time"
)

type User struct {
  FirstName string
  LastName string
  BirthDate string
  createdAt time.Time
}

func (u *User) OutputUserDetails() {
  // ...
   fmt.Println((*u).firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
  u.firtsName = ""
  u.lastName = ""
}

// constructor pattern
func NewUser(firstName, lastName, birthDate string) (*User, error) {
  if firtsName == "" || lastName == "" || birthDate == "" {
      return nil, errors.New("First name, last name, birth date is required!")
  }
  return &User {
    FirstName: firstName,
    LastName: lastName,
    BirthDate: birthDate,
    createdAt: time.Now(),
  }, nil
}
