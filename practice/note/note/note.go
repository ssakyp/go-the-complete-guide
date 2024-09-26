package note

import (
  "time"
  "encoding/json"
  "errors"
  "fmt"
  "os"
)

type Note struct {
  title string
  content string
  createdAt time.Time
}


func (note Note) Display() {
  fmt.Printf("Your not titled %v has the following content:\n\n%v", note.title, note.content)
}

func (note Note) Save() error {
  fileName := strings.ReplaceAll(note.title, " ", "_")
  fileName := strings.ToLower(fileName)

  json, err := json.Marshal(note)

  if err != nil {
    return err
  }
  return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
   if title == "" || content == "" {
    return Note{}, errors.New("Invalid input")
  }

  return Note{title, content, time.Now()}, nil
}
