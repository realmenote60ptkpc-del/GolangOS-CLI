package main

// Imports

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- GolangOS CLI v0.2 was booted --- ")

	// Scan
	for {
		fmt.Print(">: ")
		scanner.Scan()
		command := strings.ToLower(scanner.Text())

		// Commands

		switch command {
		case "help":
			fmt.Println("Command list: help, exit, about.")

		case "about":
			fmt.Println("##################")
			fmt.Println("# GolangOS CLI   #")
			fmt.Println("# Version: 0.2.1 #")
			fmt.Println("# Stage: Alpha   #")
			fmt.Println("##################")

		case "licence", "license":
			fmt.Println("# Licensed under the Fhund Free License")
            fmt.Println("-Do whatever you want, just credit the author-.")

		case "exit":
			fmt.Println("logout")
			time.Sleep(5 * time.Millisecond)
			fmt.Println("logout.")
			time.Sleep(5 * time.Millisecond)
			fmt.Println("logout..")
			time.Sleep(5 * time.Millisecond)
			fmt.Println("logout....")
			return

		default:
			fmt.Println("Command not found!")
		}
	}
}
