package repl

import (
	"bufio"
	"dynamite/src/lexer"
	"dynamite/src/logger"
	"dynamite/src/parser"
	"fmt"
	"log"
	"os"
)

func runString(s string) {
	l := lexer.New(s)
	p := parser.New(l)
	p.ParseProgram()
}

func fileMode() {
	fileName := os.Args[1]
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	str := string(b)
	runString(str)
}

var PROMT = logger.Info(">> ")

// REPL persists env values withing the same repl session
func replMode() {
	fmt.Print("\nWelcome to Dynamite REPL 👋\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		str := scanner.Text()
		runString(str)
		fmt.Println()
	}
}

func Start() {
	if(len(os.Args) > 1) {
		fileMode()
	} else {
		replMode()
	}
}