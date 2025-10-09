package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	truncated := int(math.Trunc(num))
	fmt.Printf("Truncated integer: %d\n", truncated)
}