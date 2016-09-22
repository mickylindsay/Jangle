package main

import "fmt"

func main(){
     s := "This is the Test string"
     fmt.Println([]byte(s[10:]))
     fmt.Println(s[10:])
     b := []byte(s[10:])
     fmt.Println(b)    
     
}