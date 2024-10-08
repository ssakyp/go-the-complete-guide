
package fileops

import (
  "strconv"
  "os"
  "fmt"
  "errors"
)

func GetFloatFromFile(fileName string) (float64, error) {
  data, err := os.ReadFile(fileName)

  if err != nil {
    return 1000,  errors.New("Failed to find file.") 
  }
  
  valueText := string(data)
  value, err := strconv.ParseFloat(valueText, 64)

  if err != nil {
    return 1000, errors.New("Failed to parse stored value.")
  }
  
  return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
  valueText := fmt.Sprint(balance)
  os.WriteFile(fileName, []byte(valueText), 0644)
}
