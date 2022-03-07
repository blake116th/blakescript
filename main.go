package main

import (
	"os"
	"os/user"

	"github.com/Appleby43/blakescript/repl"
)

func main() {
	_, err := user.Current()

	if err != nil {
		panic(err)
	}

	repl.Start(os.Stdin, os.Stdout)
}