package user

import (
  "fmt"
  "errors"
  "time"
)

type User struct {
  firstName string
  lastName string
  birthDate string
  createdAt time.Time
}

// go does not have inheritance, but still we can make up a new struct based on an existing one
type Admin struct {
  email string
  password string
  User
}

func (u *User) OutputUserDetails() {
  // ...
   fmt.Println((*u).firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
  u.firtsName = ""
  u.lastName = ""
}

func NewAdmin(email, pasword string) Admin {
  return Admin{
    email: email,
    password: password,
    User: User{
      firstName: "ADMIN",
      lastName: "ADMIN",
      birthDate: "---",
      createdAt: time.Now(),
    }
  }
}

// constructor pattern
func New(firstName, lastName, birthDate string) (*User, error) {
  if firtsName == "" || lastName == "" || birthDate == "" {
      return nil, errors.New("First name, last name, birth date is required!")
  }
  return &User {
    firstName: firstName,
    lastName: lastName,
    birthDate: birthDate,
    createdAt: time.Now(),
  }, nil
}
