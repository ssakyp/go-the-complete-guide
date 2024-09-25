package main

import "fmt"

func main() {
  age := 32 // regular variable

  var agePointer *int
  agePointer = &age
  
  fmt.Pritln("Age: ", agePointer) // 0x004....
  fmt.Println(*agePointer) // 32
  
  adultYears := getAdultYears(age)
  fmt.Println(adultYears)

}

func getAdultYears(age int) int {
  return age - 18
}
