package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal holds information about an animal's food, locomotion, and sound.
type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

// Eat prints the food that the animal eats.
func (a Animal) Eat() {
	fmt.Println(a.Food)
}

// Move prints the locomotion method of the animal.
func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

// Speak prints the sound the animal makes.
func (a Animal) Speak() {
	fmt.Println(a.Noise)
}

// animalMap returns a map of animal names to their Animal struct.
func animalMap() map[string]Animal {
	return map[string]Animal{
		"cow":   {Food: "grass", Locomotion: "walk", Noise: "moo"},
		"bird":  {Food: "worms", Locomotion: "fly", Noise: "peep"},
		"snake": {Food: "mice", Locomotion: "slither", Noise: "hsss"},
	}
}

// handleRequest processes a user request and prints the appropriate animal information.
func handleRequest(animalData map[string]Animal, animalName, infoType string) {
	animal, exists := animalData[strings.ToLower(animalName)]
	if !exists {
		fmt.Println("Unknown animal. Please enter cow, bird, or snake.")
		return
	}

	switch strings.ToLower(infoType) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Unknown information type. Please enter eat, move, or speak.")
	}
}

func main() {
	animalData := animalMap()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.Fields(scanner.Text())
		if len(input) != 2 {
			fmt.Println("Please enter a request in the format: <animal> <information>")
			continue
		}
		handleRequest(animalData, input[0], input[1])
	}
}
