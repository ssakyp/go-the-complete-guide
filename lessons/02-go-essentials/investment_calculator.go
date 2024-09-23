package main // every go code file needs a package
// main is the main entrance package
// must have at least one package
// the idea is to simply organize the code

// one module consists of at least one package
// go mod init ____
// go build => creates an executable file => can be executed without installing go

import (
  "math"
  "fmt"
)

const inflationRate = 2.5

func main() {
  var years float64
  var  expectedReturnRate := 5.5
  var investmentAmount float64
  
  // fmt.Print("Investment amount: ")
  outputText("Investment amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Printf("Years: ")
  fmt.Scan(&years)

  fmt.Printf("Expected return rate: ")
  fmt.Scan(&expectedReturnRate)

  futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
  // futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
  // futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

  formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
  formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
  
  //fmt.Println("Future Value: ", futureValue)
  // ` backtick for string formatting
  // fmt.Prinft(`Future Value: %.2f
  // Future Value (adjusted for Inflation): %v`, 
  // futureValue, futureRealValue)
  //fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)

  fmt.Println(formattedFV, formattedRFV)
}

func outputText(text string) {
  fmt.Print(text, text2)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64){
  fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
  rfv = futureValue / math.Pow(1 + inflationRate / 100, years)
  return fv, rfv
}
