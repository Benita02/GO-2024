package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func RunCLI() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n Choose an option:")
		fmt.Println("1. Look up a word")
		fmt.Println("2. View saved words")
		fmt.Println("3. Exit")

		fmt.Println("\nEnter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Enter the word to look up: ")
			if !scanner.Scan() {
				fmt.Println("Error reading input. Exiting...")
				return
			}
			word := scanner.Text()
			handleWordLookup(word)
		case "2":
			handleViewSavedWords()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
