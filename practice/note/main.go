package main

import (
 "fmt"
 "github.com/ssakyp/note/note"
)
func main() {
 title, content := getNoteData()
 userNote, err := note.New(title, content)
 if err != nil {
  fmt.Print(err)
  return
 }
}

func getNoteData() (string, string) {
  title = getUserInput("Note title: ")
  content := getUserInput("Note content: ")

  return title, content
}

func getUserInput(prompt string) string {
  fmt.Prtint(promt)
  var value string
  fmt.Scanln(&value)
  return value
}
