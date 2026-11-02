package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	fileList, _ := os.ReadDir("./")
	for _, file := range fileList {
		if strings.HasPrefix("output.txt", file.Name()) {
			file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()
			io.WriteString(file, content)
			return
		}
		file, err := os.Create("output.html")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		apeendContent := fmt.Sprintf("%s \n", content)
		io.WriteString(file, apeendContent)
	}
}


func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		writeToFile("output.txt", "\n"+scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func main() {
	args := os.Args[1:]
	fileExtention := getFileExt(args[0])

	if fileExtention == "txt" {
		fmt.Println("Please provide a `.MD` file...")
		return
	}
	readFile(args[0])
}
