package main

import "fmt"

const accountBalanceFile = "balance.txt"



func main() {
  var accountBalance, err = getFloatFromFile(accountBalanceFile)

  if err != nil {
    fmt.Println("ERROR")
    fmt.Println(err)
    fmt.Println("------------------------------")
    // panic("Can't continue, sorry.")
  }
  
  fmt.Println("Welcome to Go Bank!")

  for {
   presentOptions()

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
    writeFloatToFile(accountBalance, accountBalanceFile)
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
    writeFloatToFile(accountBalance, accountBalanceFile)
  default:
    fmt.Println("Goodbye!")
    fmt.Println("Thanks for choosing our bank!")
    return
}   
}



