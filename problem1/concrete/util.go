package concrete

import "fmt"

func AskForUserChoice() string {
	var choice string
	fmt.Print("\nDo you want to enter details of any other item (y/n):")
	fmt.Scan(&choice)
	return choice
}

func ValidateYesNoChoice(choice string) bool {
	if choice == "y" || choice == "n" {
		return true
	}
	return false
}
