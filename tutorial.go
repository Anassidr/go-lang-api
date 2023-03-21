package main

import "fmt"

func main() {
	fmt.Println("Let's start the quiz game!")
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome!\n", name)
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Println("I will let you play")
	} else {
		fmt.Println("naah")
		return
	}

	score := 0
	num_quest := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("correct")
		score += 1
	} else {
		fmt.Println("pff")
	}

	fmt.Printf("How many cores?")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("correct")
		score += 1
	} else {
		fmt.Println("pff")
	}

	fmt.Printf("you score %v out of %v \n", score, num_quest)
	percent := (float64(score) / float64(num_quest)) * 100
	fmt.Printf("That's %v%%", percent)
}
