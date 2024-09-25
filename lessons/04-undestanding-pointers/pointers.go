package main

import "fmt"

func main() {
  age := 32 // regular variable

  fmt.Pritln("Age: ", age)
  
  adultYears := getAdultYears(age)
  fmt.Println(adultYears)

}

func getAdultYears(age int) int {
  return age - 18
}
