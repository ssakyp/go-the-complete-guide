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

func main() {
  const inflationRate = 2.5
  var years float64
  expectedReturnRate := 5.5
  var investmentAmount float64

  fmt.Print("Investment amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Printf("Years: ")
  fmt.Scan(&years)

  fmt.Printf("Expected return rate: ")
  fmt.Scan(&expectedReturnRate)
  
  futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
  futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
  
  fmt.Println(futureValue)
  fmt.Println(futureRealValue)
}
