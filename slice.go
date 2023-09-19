package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

func main() {
	userInts := make([]int, 3, 3)

	for range userInts {
		var userInput string

		fmt.Print("Enter an Integer: ")
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == "x" {
			return
		}
		userInputInt, _ := strconv.Atoi(userInput)

		userInts[0] = userInputInt
		sort.Ints(userInts)
		fmt.Println(userInts)
	}

	for {
		var userInput string

		fmt.Print("Enter an Integer: ")
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == "x" {
			return
		}
		userInputInt, _ := strconv.Atoi(userInput)

		userInts = append(userInts, userInputInt)
		sort.Ints(userInts)
		fmt.Println(userInts)
	}
}
