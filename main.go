package main

import (
	parser "assembler-go/parser"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	countFileLines(f)
	f.Close()
}

func countFileLines(f *os.File) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		fileLine := input.Text()
		parsedInstruction, err := parser.ParseInstruction(fileLine)

		if err != nil {
			continue
		} else {
			fmt.Println(parsedInstruction)
		}
	}
}
