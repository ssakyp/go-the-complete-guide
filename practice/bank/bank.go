package main

import "fmt"

func main() {
  var accountBalance float64 = 1000
  fmt.Println("Welcome to Go Bank!")

  
  for {
  fmt.Println("What do you want to do?")
  fmt.Println("1. Check balance")
  fmt.Println("2. Deposit money")
  fmt.Println("3. Withdraw money")
  fmt.Println("4. Exit")

  var choice int
  fmt.Print("Your choice: ")
  fmt.Scan(&choice)
  fmt.Println("Your choice: ", choice)

  //wantsCheckBalance := choice == 1
  switch choice {
  case 1:
    fmt.Println("Your balance is: ", accountBalance)
  case 2:
      fmt.Println("Your deposit: ")
    var depositAmount float64
    fmt.Scan(&deposit)
    if depositAmount <= 0 {
      fmt.Println("Invalid amount. Must be greater than 0.")
      continue
    }
    // continue => skip the rest of the code
    accountBalance += depositAmount
    fmt.Println("Balance updated! New amount: ", accountBalance)
  }
  case 3: 
    var withdrawAmount float64
    fmt.Scan(&withdrawAmount)
     if withdrawAmount > 0 {
      fmt.Println("Invalid amount. You can't withdraw more than you have.")
      continue
    }
    accountBalance -= withdrawAmount
    fmt.Println("Balance updated! New amount: ", accountBalance)
  default:
    fmt.Println("Goodbye!")
    fmt.Println("Thanks for choosing our bank!")
    return
}   
 
}


