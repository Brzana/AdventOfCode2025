package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := "input.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}

	inputString := string(content)

	inputList := strings.Split(inputString, "\r\n")

	solver(inputList)
}