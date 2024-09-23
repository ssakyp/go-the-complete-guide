package main

func main() {
  var revenue float64
  var expenses float64
  var taxRate float64
  
  fmt.Print("Revenue: ")
  fmt.Scan(&revenue)

  
  fmt.Print("Expenses: ")
  fmt.Scan(&expenses)

  
  fmt.Print("Tax Rate: ")
  fmt.Scan(&taxRate)

  earningsBeforeTax := revenue - expenses
  earningsAfterTax := (revenue - expenses) / taxRate
  ratio := earningsBeforeTax / earningsAfterTax

  fmt.Println(earningsBeforeTax)
  fmt.Println(earningsAfterTax)
  fmt.Println(ratio)


  
}
