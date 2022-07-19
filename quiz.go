
package main

import "fmt"

func scoreBoard(x int) {

	fmt.Println("Thanks for completing!")

	fmt.Println("Your Score: ", x)
}

func main() {

	quizSet := map[int]string{

		1: "Who is the richest man in Africa ",

		2: "How many states of matter do we have",

		3: "How may states are in the United States",

		4: "Name of Apple's third founder",

		5: "Highest mountain in the world"}

	fmt.Print("Are you ready for the quiz? enter y(yes) n(no): ")

	var choice string

	var total int = 0

	fmt.Scanln(&choice)

	if choice != "y" {

		fmt.Println("Thanks for trying! \n we hope to see you participate next time")

	}

	if choice == "y" {

		fmt.Println(quizSet[1])

		fmt.Println("1. Warren Buffet")

		fmt.Println("2. Lucky Dare")

		fmt.Println("3. Aliko Dangote")

		fmt.Println("4. George Bush")

		fmt.Print("Answer: ")

		var answer int

		fmt.Scanln(&answer)

		if answer == 3 {

			total += 20

		}

		fmt.Println(quizSet[4])

		fmt.Println("1. Steve Jobs")

		fmt.Println("2. Steve Balmer")

		fmt.Println("3. Steve Woz")

		fmt.Println("4. Ronald Wayne")

		fmt.Print("Answer: ")

		var answer4 int

		fmt.Scanln(&answer4)

		if answer4 == 4 {

			total += 20

		}

		fmt.Println(quizSet[3])

		fmt.Println("1. 50")

		fmt.Println("2. 56")

		fmt.Println("3. 60")

		fmt.Println("4. 65")

		var answer3 int

		fmt.Scanln(&answer3)

		if answer3 == 1 {

			total += 20

		}

		fmt.Println(quizSet[2])

		fmt.Println("1. 3")

		fmt.Println("2. 5")

		fmt.Println("3. 4")

		fmt.Println("4. 2")

		var answer2 int

		fmt.Scanln(&answer2)

		if answer2 == 3 {

			total += 20

		}

		fmt.Println(quizSet[5])

		fmt.Println("1. Kilimanjaero")

		fmt.Println("2. Everest")

		fmt.Println("3. Timbuti")

		fmt.Println("4. Zion")

		var answer5 int

		fmt.Scanln(&answer5)

		if answer5 == 2 {

			total += 20

		}

	}

	scoreBoard(total)

}