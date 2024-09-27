package main

func main() {
  userNames := make([]string, 2)

  userNames = append(userNames, "Max")
  userNames = append(userNames, "Manuel")

  fmt.Println(userNames)
}
