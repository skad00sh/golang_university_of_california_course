package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}

func main(){
	fmt.Printf("\nEnter the file name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	People := make([]Person, 0)
	fd, _ := os.Open(fileName)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, " ")
		localPerson := Person{firstName:names[0], lastName:names[1]}
		People = append(People, localPerson)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nIterating through struct:\n\n")
	for _, v := range People{
		fmt.Printf("First Name :  %s\nLast Name  :  %s\n***********\n", v.firstName, v.lastName)
	}
}
