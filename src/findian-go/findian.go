package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		lower := strings.ToLower(input)
		if strings.HasPrefix(lower, "i") &&
			strings.HasSuffix(lower, "n") &&
			strings.Contains(lower, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
