package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func BubbleSort(integersSlice *[]int) {
	sliceLen := len(*integersSlice)
	isSorted := false

	for !isSorted {
		isSorted = true

		for index, integer := range *integersSlice {

			// Skip the last element as it has no further elements to be compared with
			if index + 1 == sliceLen {
				continue
			}

			// If current element is greater than the next one
			if (*integersSlice)[index + 1] < integer {
				Swap(integersSlice, index)

				isSorted = false
			}
		}
	}
}

func Swap(integersSlice *[]int, index int) {
	swap := (*integersSlice)[index]
	(*integersSlice)[index] = (*integersSlice)[index + 1]
	(*integersSlice)[index + 1] = swap
}

func main() {
	fmt.Print("Enter 10 integers space seperated (2 3 55...): ")
	in := bufio.NewReader(os.Stdin)
	rawUserInput, _, _ := in.ReadLine()
	strings := strings.Split(string(rawUserInput), " ")

	if len(strings) != 10 {
		fmt.Println("Please enter 10 integers!")
		return
	}

	var integers []int
	for i := 0; i < 10; i++ {
		integer, _ := strconv.Atoi(strings[i])
		integers = append(integers, integer)
	}

	BubbleSort(&integers)
	fmt.Println("\nSorted using Bubble Sort")
	for _, integer := range integers {
		fmt.Print(integer)
		fmt.Print(" ")
	}
	fmt.Println("")
}
