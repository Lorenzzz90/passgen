package main

import (
	"fmt"
	"os"
)

func main() {
	var characters string = ""
	if len(os.Args) > 2 {
		// FIXME: if you give a wrong number of arguments, it continues and reports a panic
		// immediately stop the program as soon as you hit a blocker for your program.
		fmt.Println("You should provide zero or one command line arguments.")
	} else if len(os.Args) == 1 {
		characters = "abcdefgh"
	} else {
		characters = os.Args[1]
	}
	for i := 0; len(characters) < 8; i++ {
		char := characters[i]
		characters = characters + string(char)
	}

	// FIXME: print a better message. Not only a password. You have to write meaningful UI (the user should not guess what you're doing!)
	for i := 0; i < 8; i++ {
		fmt.Printf("%c", characters[i])
	}
}
