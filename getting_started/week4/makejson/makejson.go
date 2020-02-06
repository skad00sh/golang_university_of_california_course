package main

import (
	"encoding/json"
	"fmt"
)

func main(){
  var name_ip string
  var addr_ip string

	fmt.Printf("\nEnter your name: ")
  fmt.Scan(&name_ip)
	fmt.Printf("\nEnter your address: ")
  fmt.Scan(&addr_ip)

	Person := make(map[string] string)
	Person["name"] = name_ip
	Person["address"] = addr_ip

	json_op, _ := json.Marshal(Person)

	fmt.Printf("\nPerson's details in a JSON structure:\n\n %s\n\n", json_op)

}
