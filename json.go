package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	payload := make(map[string]string)
	var userInput string

	fmt.Print("Enter name: ")
	fmt.Scan(&userInput)
	payload["name"] = userInput

	fmt.Print("Enter address: ")
	fmt.Scan(&userInput)
	payload["address"] = userInput

	json, _ := json.Marshal(payload)
	fmt.Println(json)
}
