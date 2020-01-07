package main

import "fmt"

func main()  {
    var floatInput float64
    fmt.Printf("Enter your float number:")
    fmt.Scan(&floatInput)

    var intOutput int64 = int64(floatInput)
    fmt.Printf("Integer part of entered float is: %d",intOutput)
}
