package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "strings"
)

func indexOf(item string) int {
	for i, char := range item {
		if char == '.' {
			return i
		}
	}
	return -1
}

func getFileExt(filename string) string {
	index := indexOf(filename)
	if index == -1 {
		return ""
	}

	return filename[index+1:]
}

func writeToFile(filename string, content string) {
	// os.O_APPEND: Appends data to the file if it exists
	// os.O_CREATE: Automatically creates the file if it does not exist
	// os.O_WRONLY: Opens the file in write-only mode
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening/creating file:", err)
		return
	}
	defer file.Close() // Safely closes when the function returns

	// Write the content followed by a clean newline character
	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func readFile(filename string, targetFilename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// Pass the target filename dynamically instead of hardcoding it
		writeToFile(targetFilename, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func main() {
	args := os.Args[1:]
	fileExtention := getFileExt(args[0])

	if fileExtention != "md" {
		fmt.Println("Please provide a `.MD` file...")
		return
	}
	readFile(args[0], args[1])
}
