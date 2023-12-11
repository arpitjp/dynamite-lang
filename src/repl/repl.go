package repl

import (
	"bufio"
	"dynamite/src/lexer"
	"dynamite/src/logger"
	"dynamite/src/tokens"
	"fmt"
	"log"
	"os"
)

func runString(s string) {
	l := lexer.New(s)
	for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
		tok.Inspect()
	}
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
	fmt.Println("Welcome to Dynamite REPL 👋")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMT)
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