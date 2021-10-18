package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	userInput := readString("What is your name?")

	fmt.Println("Your name is", userInput)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
