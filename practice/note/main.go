package main

import (
 "fmt"
 "github.com/ssakyp/note/note"
 "bufio"
 "os"
 "strings"
)
func main() {
 title, content := getNoteData()
 userNote, err := note.New(title, content)
 if err != nil {
  fmt.Print(err)
  return
 }
 userNote.Display()
}

func getNoteData() (string, string) {
  title = getUserInput("Note title: ")
  content := getUserInput("Note content: ")

  return title, content
}

func getUserInput(prompt string) string {
  fmt.Prtint(promt)
  // scan cannot work for multiple space lines
  // var value string
  // fmt.Scanln(&value)
  reader := bufio.NewReader(os.Stdin)
  text, err := reader.ReadString('\n')
  if err != nil {
   return ""
  }
 
  text = strings.TrimSuffix(text, "\n")
  text = strings.TrimSuffix(text, "r")
 
  return text
}
