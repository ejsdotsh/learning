// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ejsdotsh/learning/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! this is the Monkey programming language.\n", user.Username)
	fmt.Printf("type some commands.\n")
	// start the REPL
	repl.Start(os.Stdin, os.Stdout)
}
