package main

// Imports

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- GolangOS CLI v0.1 was booted --- ")

	// Scan
	for {
		fmt.Print(">: ")
		scanner.Scan()
		command := strings.ToLower(scanner.Text())

		// Commands

		if command == "help" {
			fmt.Println("Список команд: help, exit, about, golangos --version.")
		} else if command == "golangos --version" {
			fmt.Println("##-- GolangOS Alpha 0.1 --##")
		} else if command == "about" {
			fmt.Println("################")
			fmt.Println("# GolangOS CLI #")
			fmt.Println("# Version: 0.1 #")
			fmt.Println("# Stage: Alpha #")
			fmt.Println("################")
		} else if command == "exit" {
			fmt.Println("Выход из GolangOS CLI 0.1")
			break
		} else {
			fmt.Println("Command not found!")
		}
	}
}
