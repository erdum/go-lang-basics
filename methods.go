package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Cow struct {
	food string
	locomotion string
	noise string
}

type Bird struct {
	food string
	locomotion string
	noise string
}

type Snake struct {
	food string
	locomotion string
	noise string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func getAnimalInfo(animal Animal, action string) {

	switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
	}
}

func main() {
	cowStorage := make(map[string]Cow)
	birdStorage := make(map[string]Bird)
	snakeStorage := make(map[string]Snake)

	for {
		fmt.Print(">");
		in := bufio.NewReader(os.Stdin)
		rawUserInput, _, _ := in.ReadLine()
		input := strings.Split(string(rawUserInput), " ")
		command := strings.ToLower(input[0])
		firstArg := input[1]
		secondArg := strings.ToLower(input[2])

		if command == "newanimal" {

			switch secondArg {
				case "cow":
					cowStorage[firstArg] = Cow{
						food: "grass",
						locomotion: "walk",
						noise: "moo",
					}
				case "bird":
					birdStorage[firstArg] = Bird{
						food: "worms",
						locomotion: "fly",
						noise: "peep",
					}
				case "snake":
					snakeStorage[firstArg] = Snake{
						food: "mice",
						locomotion: "slither",
						noise: "hsss",
					}
			}
			fmt.Println("Created it!\n")
		} else if command == "query" {
			cow, isCow := cowStorage[firstArg]
			bird, isBird := birdStorage[firstArg]
			snake, isSnake := snakeStorage[firstArg]

			if isCow {
				getAnimalInfo(cow, secondArg)
			} else if isBird {
				getAnimalInfo(bird, secondArg)
			} else if isSnake {
				getAnimalInfo(snake, secondArg)
			} else {
				fmt.Println("The animal " + firstArg + " you specified is not present!\n")
				continue
			}
		}
	}
}
