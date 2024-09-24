package main

import (
  "fmt"
  "errors"
  "os"
)

var earningsAmountFile = "earnings.txt"

func main() {
  // var revenue float64
  // var expenses float64
  // var taxRate float64
  // fmt.Print("Revenue: ")
  // fmt.Scan(&revenue)
  for {
  revenue, err := prompt("Revenue")
  if err != nil {
    fmt.Print(err)
    continue
  }
  
  // fmt.Print("Expenses: ")
  // fmt.Scan(&expenses)
  expenses, err := prompt("Expenses")
   if err != nil {
    fmt.Print(err)
    continue
  }
  
  // fmt.Print("Tax Rate: ")
  // fmt.Scan(&taxRate)
  taxRate, err := prompt("Tax Rate")
    if err != nil {
    fmt.Print(err)
    continue
  }

  // earningsBeforeTax := revenue - expenses
  // earningsAfterTax := earningsBeforeTax * (1 - taxRate / 100)
  // ratio := earningsBeforeTax / earningsAfterTax
  
  earningsBeforeTax, earningsAfterTax, ratio := profitCalculator(revenue, expenses, taxRate)
  fmt.Printf("%.1f\n", earningsBeforeTax)
  fmt.Printf("%.1f\n", earningsAfterTax)
  fmt.Printf("%.3f\n", ratio)
    return
}
}

func prompt(text string) (float64, error) {
  var num float64
  fmt.Print(text, ": ")
  fmt.Scan(&num)
  if num <= 0 {
    return 0, errors.New("Invalid input. Negative amount or 0 entered.")
  }
  return num, nil
}

func profitCalculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
  ebf := revenue - expenses
  profit := ebf * (1 - taxRate / 100)
  ratio := ebf / profit

  text := fmt.Sprintf("Earnings Before tax: %.1f\nEarnings After tax or profit: %.1f\nRatio: %.1f\n", ebf, profit, ratio)
  os.WriteFile(earningsAmountFile, []bye(text), 0644)
  return ebf, profit, ratio
}
