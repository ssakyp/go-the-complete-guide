package main

import "fmt"

func main() {
  age := 32 // regular variable

  var agePointer *int
  agePointer = &age
  
  fmt.Pritln("Age: ", agePointer) // 0x004....
  fmt.Println(*agePointer) // 32
  
  editAgeToAdultYears(&age)
  fmt.Println(age) // 14

}

func editAgeToAdultYears(age *int) {
  // no copy is created
  *age -= 18
  // we can't perform pointer arithmetics
}
