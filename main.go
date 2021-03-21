package main

import (
	"fmt"
	"os"
	"os/user"
	"mymarkdown/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s. This is mymarkdown\n", user.Username)
	fmt.Printf("Type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
