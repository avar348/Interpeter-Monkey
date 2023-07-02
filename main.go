package main

import (
	"fmt"
	"learningmonky/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey Programming Language!", user.Username)
	fmt.Println("Feel free to type in any line")
	repl.Start(os.Stdin, os.Stdout)

}
