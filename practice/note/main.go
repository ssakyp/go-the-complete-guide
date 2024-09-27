package main

import (
 "fmt"
 "github.com/ssakyp/note/note"
 "github.com/ssakyp/note/todo"
 "bufio"
 "os"
 "strings"
)

type saver interface {
 Save() error
}

func main() {
 title, content := getNoteData()
 todoText := getUserInput("Todo text: ")

 todo, err := todo.New(todoText)

  if err != nil {
  fmt.Print(err)
  return
 }
 
 userNote, err := note.New(title, content)
 if err != nil {
  fmt.Print(err)
  return
 }
 
 todo.Display()
 err = saveData(todo)
  if err != nil {
  return
 }
 
 userNote.Display()
 err = saveData(userNote)
 if err != nil {
  return
 }
}


func saveData(data saver) {
 err := data.Save()
 if err != nil {
  fmt.Println("Saving the note failed.")
  return err
 }

 fmt.Println("Saving the note succeeded!")
 return nil
}



func getNoteData() (string, string) {
  title = getUserInput("Note title: ")
  content := getUserInput("Note content: ")

  return title, content
}

func getUserInput(prompt string) string {
  fmt.Prtintf("%v ", promt)
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
