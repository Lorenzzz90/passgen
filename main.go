package main

import "fmt"
import "os"

func main() {

	var characters string = ""
	if len(os.Args) > 2 {
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

	for i := 0; i < 8; i++ {
		fmt.Printf("%c", characters[i])
	}
}
