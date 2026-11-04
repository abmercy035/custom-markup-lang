package main

import (
	"bufio"
	"fmt"
	"os"
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

func TransitFile(filename string, targetFilename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Open destination once

	dst, err := os.OpenFile(targetFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(dst)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Println(scanner.Text())
		writer.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	writer.Flush() // Ensure all buffered data is written to the destination file
}

func main() {
	args := os.Args[1:]
	fileExtention := getFileExt(args[0])

	if fileExtention != "md" {
		fmt.Println("Please provide a `.MD` file...")
		return
	}
	TransitFile(args[0], args[1])
}
