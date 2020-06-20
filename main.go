package main

import (
	"bufio"
	"github.com/varshard/monkey/repl"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	repl.Run(io)
}
