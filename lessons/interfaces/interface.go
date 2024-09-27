package main

func main() {
  
}

func add(a, b interface{}) interface{} {
  aInt, aIsInt := a.(int)
  bInt, bIsInt := b.(int)

  if aIsInt && bIsInt {
    return aInt + bInt
  }

   aFloat, aIsFloat := a.(float64)
  bFloat, bIsFloat := b.(float64)

  if aIsFloat && bIsFloat {
    return aFloat + bFloat
  }

   aStr, aIsStr := a.(string)
  bStr, bIsStr := b.(string)

  if aIsStr && bIsStr {
    return aStr + bStr
  }

  
  
}
