package main

import(
	"fmt"
	"math"
)

func GenDisplaceFn(acc float64, vI float64, sI float64) func (float64) float64 {
	return func (time float64) float64 {
		return ((acc * math.Pow(time, 2)) / 2) + (vI * time) + sI
	}
}

func main() {
	var acc, vI, sI, time float64

	fmt.Print("Enter the acceleration: ")
	fmt.Scan(&acc)

	fmt.Print("Enter the initial velocity: ")
	fmt.Scan(&vI)

	fmt.Print("Enter the initial displacment: ")
	fmt.Scan(&sI)

	fmt.Print("Enter the travel time: ")
	fmt.Scan(&time)

	displace := GenDisplaceFn(acc, vI, sI)

	fmt.Print(displace(time))
	fmt.Println(" meters.")
}
