package main

func main() {
 title, content, err := getNoteData()

  if err != nil {
    fmt.Println(err)
    return
  }
}

func getNoteData() (string, string, error) {
  title, err := getUserInput("Note title: ")

  if err != nil {
    fmt.Println(err)
    return "", "", err
  }

  content, err := getUserInput("Note content: ")

  if err != nil {
    fmt.Println(err)
     return "", "", err
  }

  return title, content, nil
}

func getUserInput(prompt string) (string, error) {
  fmt.Prtint(promt)
  var value string
  fmt.Scanln(&value)

  if value == "" {
    return "", errors.New("Invalid input")
  }

  return value, nil
}
