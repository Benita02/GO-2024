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
	}
}
