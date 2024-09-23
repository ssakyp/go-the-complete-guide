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
  var investmentAmount = 1000
  var expectedReturnRate = 5.5
  var years = 10

  var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
  fmt.Println(futureValue)
}
