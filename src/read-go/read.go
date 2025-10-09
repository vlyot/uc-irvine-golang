package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct with fixed-size string fields (up to 20 chars)
type Name struct {
	Fname string
	Lname string
}

func main() {
	// Prompt user for filename
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue // skip lines that don't have exactly two fields
		}
		fname := truncate(parts[0], 20)
		lname := truncate(parts[1], 20)
		names = append(names, Name{Fname: fname, Lname: lname})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the names
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.Fname, n.Lname)
	}
}

// truncate ensures the string is at most 20 characters
func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	return s
}
