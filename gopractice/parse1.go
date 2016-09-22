package main

import "strings"
import "fmt"
import "bufio"
import "os"

func finder (test string) {
	login := "0000"
	message := "0001"
	command := "0002"

	if(strings.Contains(test,login)) {
		fmt.Print("login\n")
	}
	if(strings.Contains(test,message)) {
		fmt.Print("message\n")
	}
	if(strings.Contains(test,command)) {
		fmt.Print("command\n")
	}
}

func main() {
	fmt.Print("Enter text: ")
    reader := bufio.NewReader(os.Stdin)
    text,_:= reader.ReadString('\n')
    
	arr := []byte(text)
	str := string(arr[:])
	finder(str)
	fmt.Println(text[4:])
	fmt.Println([]byte(text[4:]))
}