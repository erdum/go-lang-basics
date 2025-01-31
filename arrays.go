package main

import (
	"fmt"
	"reflect"
)

func main() {
	// This is an array only declared
	// var ints [5]int

	// This is an array with type and initial values
	// var ints_two = [5]int{0, 1, 2, 3, 4}

	// This is an array with inferred length and initial values
	// var ints_three = [...]int{0, 5, 5, 6}

	// ints_four := [2]int{3, 5}
	// ints_five := [...]int{5, 7}

	// Slices
	// ints_six := make([]int, 2)
	// ints_seven := make([]int, 2, 5)

	// Create a slice from an array
	// ints_eight := ints_five[:]
	// ints_nine := ints_five[1:]
	// ints_ten := ints_five[:2]

	// Remove item from a slice
	// test := append(ints_two[:3], ints_two[4:]...)

	// A slice of slices of string
	multi := [][]string{{"lorem", "ipsum"}, {"doler", "sit"}}

	fmt.Println(reflect.TypeOf(multi))
	fmt.Println(multi)
}
