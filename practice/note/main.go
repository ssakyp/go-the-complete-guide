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

// type displayer interface {
//  Display()
// }

// embedding
type outputtable interface {
 saver
 Display()
}


// type outputtable interface {
//  Save() error
//  Display
// }

func main() {
 printSomething(1)
 printSomething(2.5)
 printSomething("Any type")


 
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

 err = outputData(todo)
 if err != nil {
  return
 }
 
 outputData(userNote)
}

// interface {} any type 
func printSomething(value interface{}) {
 intVal, ok := value.(int) // will be true if val is int

 if ok {
  fmt.Println("Integer: ", intVal)
  return
 }

  floatVal, ok := value.(float64) // will be true if val is float

 if ok {
  fmt.Println("Float: ", floatVal)
  return
 }

  stringVal, ok := value.(string) // will be true if val is string

 if ok {
  fmt.Println("String: ", stringVal)
  return
 }


 
 // switch value.(type) {
 //  case int:
 //   fmt.Println("Integer: ", value)
 //  case float64:
 //   fmt.Println("Float: ", value)
 //  case string:
 //   fmt.Println(value)
 // }
}

func outputData(data outputtable) error {
 data.Display()
 return saveData(data)
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
