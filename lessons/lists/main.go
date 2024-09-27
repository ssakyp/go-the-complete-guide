package main

// type aliases
type floatMap map[string]float64

func (m floatMap) output() {
  fmt.Println(m)
}

func main() {
  // if we know the number of elements in advance use make
  userNames := make([]string, 2, 5)
  //userNames := []string{}
  
  userNames[0] = "Julie"
  userNames = append(userNames, "Max")
  userNames = append(userNames, "Manuel")

  fmt.Println(userNames)

  // for maps only len can be set
  courseRatings := make(floatMap, 3)
  //courseRatings := map[string]float64{}

  courseRatings["go"] = 4.7
  courseRatings["reac"] = 4.8
  courseRatings["node"] = 4.7

  fmt.Println(courseRatings)

  for index, value := range userNames {
    fmt.Println(index)
    fmt.Println(value)
  }
}
