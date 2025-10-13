package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface defines the behavior of all animals
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type implements Animal interface
type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Bird type implements Animal interface
type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Snake type implements Animal interface
type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

// animalRegistry stores created animals by name
var animalRegistry = make(map[string]Animal)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.Fields(scanner.Text())
		if len(input) != 3 {
			fmt.Println("Invalid command format. Please enter three words.")
			continue
		}
		command, name, param := input[0], input[1], input[2]
		switch command {
		case "newanimal":
			createAnimal(name, param)
		case "query":
			queryAnimal(name, param)
		default:
			fmt.Println("Unknown command. Use 'newanimal' or 'query'.")
		}
	}
}

// createAnimal creates a new animal and adds it to the registry
func createAnimal(name, animalType string) {
	var animal Animal
	switch strings.ToLower(animalType) {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Println("Unknown animal type. Use 'cow', 'bird', or 'snake'.")
		return
	}
	animalRegistry[name] = animal
	fmt.Println("Created it!")
}

// queryAnimal prints requested information about the animal
func queryAnimal(name, info string) {
	animal, exists := animalRegistry[name]
	if !exists {
		fmt.Println("Animal not found.")
		return
	}
	switch strings.ToLower(info) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Unknown query. Use 'eat', 'move', or 'speak'.")
	}
}
