package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	testSlice := buildAppendString()

	fmt.Println(testSlice)

}

func buildAppendString() []string {
	var combinedEntry []string

	//Enter URL
	url := getUserInput("Enter URL: ")
	combinedEntry = append(combinedEntry, strings.TrimSpace(url))

	//Enter RITM
	ritm := getUserInput("Enter RITM: ")
	combinedEntry = append(combinedEntry, strings.TrimSpace(ritm))

	//Enter Date
	date := getUserInput("Enter Date (Format DDMMYYYY): ")
	for len(date) < 9 || len(date) > 9 {
		fmt.Println("Invalid Date, Please Re-Enter, current date is length", len(date))
		date = getUserInput("Enter Date (Format DDMMYYY): ")
	}

	combinedEntry = append(combinedEntry, strings.TrimSpace(date))

	return combinedEntry
}

func getUserInput(inputText string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputText)
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	return text
}
