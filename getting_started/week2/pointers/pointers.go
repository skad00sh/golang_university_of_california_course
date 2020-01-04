package main

import "fmt"
func  main()  {
  var x int = 1
  var y int
  var ip *int
  fmt.Println("Value of x is: ",x)
  fmt.Println("Value of y is :", y)
  fmt.Println("Data at memory address of int is: ", ip)
  ip = &x
  y = *ip
  fmt.Println("Memory address(pointer) of x is: ", ip) //memory address of x
  fmt.Println("Data at memory address of x is: ", y) //data at ip
  a := x
  fmt.Println("Memory address of a is: ", &a)
  a_pointer := &a
  fmt.Println("Data at memory address of a is: ", *a_pointer)

  ptr := new(int)
  fmt.Println("Value of ptr is: ", ptr)
  fmt.Println("Memory address of ptr is: ", &ptr)
  fmt.Println("Data at memory location of ptr is: ", *ptr) 
}

//Output
/*
Value of x is:  1
Value of y is : 0
Data at memory address of int is:  <nil>
Memory address(pointer) of x is:  0xc000096000
Data at memory address of x is:  1
Memory address of a is:  0xc000096028
Data at memory address of a is:  1
*/
