package main

import "strings"
import "fmt"

func finder (test string) {
	login := "0000"
	message := "0001"
	command := "0002"

	if(strings.Contains(test,login)) {
		fmt.Print("login")
	}
	if(strings.Contains(test,message)) {
		fmt.Print("message")
	}
	if(strings.Contains(test,command)) {
		fmt.Print("command")
	}
}

func main() {
	reader := ""
	fmt.Print("Enter text: ")
	fmt.Scanln(&reader)
	finder(reader)
}