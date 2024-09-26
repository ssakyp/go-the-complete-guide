package note

import (
  "time"
  "encoding/json"
  "errors"
  "fmt"
  "os"
)

type Note struct {
  Title string
  Content string
  CreatedAt time.Time
}


func (note Note) Display() {
  fmt.Printf("Your not titled %v has the following content:\n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
  fileName := strings.ReplaceAll(note.title, " ", "_")
  fileName := strings.ToLower(fileName) + ".json"

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
