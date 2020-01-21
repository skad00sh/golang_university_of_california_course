package main

import (
	"fmt"
	"sort"
	"strconv"
)

func slc() {
	fmt.Printf("\nEnter number whenever prompt. Press x or X or Ctrl + X to exit.\n")
	stop := false
	var s string
	sl := make([]int, 0, 3)
	fmt.Printf("\nSlice initialized with length of %d and capacity of %d.\n	Initialized slice = %d\n", len(sl), cap(sl), sl)
	sl = append(sl, 0, 0, 0)
	count := 0
	for stop == false {
		fmt.Printf("\nEnter an integer to be added to the slice : ")
		_, _ = fmt.Scan(&s)
		if s == "X" || s == "x" {
			stop = true
		} else {
			sInt, _ := strconv.Atoi(s)
			if count < 3 {
				sl[count] = sInt
				count++
				sort.Ints(sl[0:count])
				fmt.Printf("\nSlice sorted:  ")
				fmt.Printf("%d\n", sl)

			} else {
				sl = append(sl, sInt)
				count++
				sort.Ints(sl)
				fmt.Printf("\nSlice sorted:  ")
				fmt.Printf("%d", sl)

			}
		}
	}

}

func main() {
	slc()
}
