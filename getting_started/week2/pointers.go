package main

import "fmt"
func  main()  {
  var x int = 1
  var y int
  var ip *int
  fmt.Println("Value of x is: ",x)
  fmt.Println("Value of y is : ",y)
  fmt.Println("memory location(pointer) of int is: ", ip)

  ip = &x
  y = *ip
  fmt.Println("Vaalue of ip is: ",ip)
  fmt.Println("Memory location of ip is: ", y)
  fmt.Println("Memory location of x is: ", *x)
}
