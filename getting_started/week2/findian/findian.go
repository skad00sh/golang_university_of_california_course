package main

import ("fmt"
        "strings")

func main()  {
  var y string
  fmt.Printf("Enter the string: ")
  fmt.Scan(&y)
  var x string = strings.ToLower(y)

  if strings.HasPrefix(x, "i") && strings.ContainsAny(x, "a") && strings.HasSuffix(x, "n") {
    fmt.Printf("Found!")
  } else {
    fmt.Printf("Not Found!")
  }
}
