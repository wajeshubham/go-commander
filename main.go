package main

import (
	"os"

	"github.com/wajeshubham/commander/package/commander"
)

func main() {
	commander.ExecuteCommand("$ ", os.Stdin, os.Stdout)
}
