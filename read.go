package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter file name to be read: ")
	in := bufio.NewReader(os.Stdin)
	rawUserInput, _, _ := in.ReadLine()

	fileName := string(rawUserInput)
	fileBytes, e := os.ReadFile(fileName)

	if e != nil {
		fmt.Println("Error while reading file " + fileName)
		return
	}

	fileText := string(fileBytes)
	lines := strings.Split(fileText, "\n")

	var namesSlice []*Name

	for _, line := range lines {

		// To make sure empty lines & lines with missin fname or lname does not make it through
		if len(line) > 3 {
			names := strings.Split(line, " ")
			fname := names[0]
			lname := names[1]

			name := new(Name)
			name.fname = fname
			name.lname = lname

			namesSlice = append(namesSlice, name)
		}
	}

	for _, nameStruct := range namesSlice {
		test := *nameStruct
		fmt.Println(test.fname, test.lname)
	}
}
