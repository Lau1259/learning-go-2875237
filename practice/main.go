package main

import (
	"fmt"
)

const first = "Lautaro"

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	fmt.Println("")
	fmt.Printf("Hi %s! glad to see you learning GO!", first)

}
