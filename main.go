package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey there, %s! Welcome to Monkey.\n", user.Username)
	fmt.Printf("Type in the commands of your liking below...\n")
	repl.Start(os.Stdin, os.Stdout)
}
