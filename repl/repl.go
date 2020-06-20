package repl

import (
	"bufio"
	"fmt"
	"github.com/varshard/monkey/eval"
	"github.com/varshard/monkey/parser"
	"io"
)

func Run(in io.Reader) {
	scanner := bufio.NewScanner(in)
	fmt.Println("Welcome to monkey REPL:")
	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(line)
		program := p.ParseProgram()
		if len(p.Errors) > 0 {
			printErrors(p.Errors)
			continue
		}

		obj := eval.Eval(program)
		fmt.Printf("out: %s\n", obj.String())
	}
}

func printErrors(errs []error) {
	fmt.Println("Errors:")
	for _, err := range errs {
		fmt.Printf("%s\n", err.Error())
	}
}
