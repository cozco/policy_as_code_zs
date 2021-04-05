package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var combinedEntry []string

	url := getUserInputTest("Enter URL: ")
	s := strings.TrimSpace(url)
	combinedEntry = append(combinedEntry, s)
	fmt.Println(combinedEntry)
}

func getUserInputTest(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputText)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	return text
}
