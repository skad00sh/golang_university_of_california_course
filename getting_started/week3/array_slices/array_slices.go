package main

import (
	"fmt"
)

func main() {
	x := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	s1 := x[1:3]
	s2 := x[2:5]

	fmt.Println(s1[1], s2[0])

	x[2] = 99

	fmt.Println(s1[1], s2[0])
	
	fmt.Println(len(s1), cap(s2))

}
