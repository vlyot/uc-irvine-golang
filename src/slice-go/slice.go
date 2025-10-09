package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	slice := make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer or 'X' to quit: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.EqualFold(input, "X") {
			fmt.Println("Exiting program.")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X'.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)
		fmt.Println("Sorted slice:", slice)
	}
}
