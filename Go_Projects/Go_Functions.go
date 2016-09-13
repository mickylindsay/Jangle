//Test function calls, and returning various functions

package main

import (
	"fmt"
	"strings"
	"strconv"
	"bytes"
)

func test( x, y int) (int, string){

	var z string
	copy(z, strconv.Itoa(x))
	copy (z,strconv.Itoa(y))
	
	return x+y,z

}

func main(){
	a, b := test(4,5)
	copy(b, " ") 
	copy(b, a)
	fmt.Println(b)

}