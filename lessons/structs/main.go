package main

type str string

func (text str) log() {
  fmt.Println(text)
}

func main() {
  var name str

  name = "Max"
  name.log()
}
