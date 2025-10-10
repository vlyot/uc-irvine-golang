package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GenDisplaceFn returns a function that calculates displacement based on time t
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}

// promptFloat prompts the user for a float64 value with the given message
func promptFloat(promptMessage string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(promptMessage)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

func main() {
	acceleration := promptFloat("Enter acceleration (a): ")
	initialVelocity := promptFloat("Enter initial velocity (vo): ")
	initialDisplacement := promptFloat("Enter initial displacement (so): ")

	displacementFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	time := promptFloat("Enter time (t): ")
	displacement := displacementFn(time)

	fmt.Printf("Displacement after %.2f seconds is %.4f\n", time, displacement)
}
