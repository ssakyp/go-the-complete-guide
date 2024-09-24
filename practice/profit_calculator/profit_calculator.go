package main

func main() {
  var revenue float64
  var expenses float64
  var taxRate float64
  
  // fmt.Print("Revenue: ")
  // fmt.Scan(&revenue)
  prompt("Revenue", &revenue)
  
  // fmt.Print("Expenses: ")
  // fmt.Scan(&expenses)
  prompt("Expenses", &expenses)
  
  // fmt.Print("Tax Rate: ")
  // fmt.Scan(&taxRate)
  prompt("Tax Rate", &taxRate)

  // earningsBeforeTax := revenue - expenses
  // earningsAfterTax := earningsBeforeTax * (1 - taxRate / 100)
  // ratio := earningsBeforeTax / earningsAfterTax
  
  earningsBeforeTax, earningsAfterTax, ratio := profitCalculator(revenue, expenses, taxRate)
  fmt.Printf("%.1f\n", earningsBeforeTax)
  fmt.Printf("%.1f\n", earningsAfterTax)
  fmt.Printf("%.3f\n", ratio)


}

func prompt(a string, b *float64) {
  fmt.Print(a,": ")
  fmt.Scan(b)
}

func profitCalculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
  ebf := revenue - expenses
  profit := ebf * (1 - taxRate / 100)
  ratio := ebf / profit

  return ebf, profit, ratio
}
