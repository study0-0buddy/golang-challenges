package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string = "log.txt"

func main() {
	// open file or create
	file := openFile(fileName)

	// new scanner to read from console
	scanner := bufio.NewScanner(os.Stdin)
	// loop to listen for user's input
	for {
		fmt.Print("Enter text: ")
		// scan a line
		scanner.Scan()
		// get scanned text
		text := scanner.Text()
		// if user's text is quit, finish the program
		if strings.ToLower(text) == "quit" {
			fmt.Println("quitting..")
			break
		}
		// write scanned text to the file
		if _, err := file.WriteString(text + "\n"); err != nil {
			// if error, log the message and exit with code 1
			log.Fatalf("failed to write to file. Error: %s",
				err.Error())
		}
	}
}

// open file or create and open if doesn't exist
func openFile(fileName string) *os.File {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if error, log the message and exit with code 1
	if err != nil {
		log.Fatalf("failed to create/open file %s. Error: %s",
			fileName, err.Error())
	}
	return f
}
