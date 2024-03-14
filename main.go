package main

import (
	"fmt"

	"github.com/MarkTBSS/EP4-CH6-Init_Function/users" // Importing the user package
)

func init() {
	fmt.Println("init from main package [2]")
}

func main() {
	fmt.Println("main from main package [3]")
	users.Main() // Calling the main function from the user package
}
