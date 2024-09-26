package note

import "time"

type Note struct {
  title string
  content string
  createdAt time.Time
}

func New(title, content string) Note {
  return Note{title, content, time.Now()}
}
