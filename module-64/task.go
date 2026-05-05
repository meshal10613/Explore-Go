package main

import "fmt"

func task() {
	//? Problem Solving
	displayMenu := func() {
		fmt.Println("Welcome to grade calculator")
		fmt.Println("(1) Calculate grade")
		fmt.Println("(2) Check Pass/Fail Status")
		fmt.Println("Exit the program")
		fmt.Print("Choose an option: ")
	}

	var choice int
	running := true

	for running {
		displayMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter a score (0 - 100): ")
			var score int
			fmt.Scan(&score)
			result := calculateGrade(score)

			if result == "Invalid score!" {
				fmt.Println("Choose a valid number from 0 to 100")
				running = false
				return
			}
			fmt.Println("You've got", result)
			running = false
		case 2:
			fmt.Print("Enter your score: ")
			var score int
			fmt.Scan(&score)
			result := checkPassFail(score)

			if result == "Invalid score!" {
				fmt.Println("Choose a valid number from 0 to 100")
				running = false
				return
			}
			fmt.Println("You've ", result)
			running = false
		case 3:
			fmt.Println("Exiting the program...")
			running = false
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func calculateGrade(score int) string {
	if score >= 80 && score <= 100 {
		return "A+"
	} else if score >= 70 && score < 80 {
		return "A"
	} else if score >= 60 && score < 70 {
		return "A-"
	} else if score >= 50 && score < 60 {
		return "B"
	} else if score >= 40 && score < 50 {
		return "C"
	} else if score <= 33 && score < 40 {
		return "D"
	} else if score <= 0 && score < 33 {
		return "F"
	} else {
		return "Invalid score!"
	}
}

func checkPassFail(score int) string {
	switch {
	case score >= 33 && score <= 100:
		return "Pass"
	case score <= 0 && score < 33:
		return "Fail"
	default:
		return "Invalid score!"
	}
}
