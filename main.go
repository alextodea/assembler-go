package main

import (
	parser "assembler-go/parser"
	pipelines "assembler-go/pipelines"
	symbolshandler "assembler-go/symbolsHandler"
	"fmt"
	"os"
)

func main() {

	assebmlyInstructionsWithLabels, err := pipelines.PreProcessData(os.Args[1])
	check(err)
	assemblyInstructionsWithoutLabels, err := symbolshandler.HandleSymbols(assebmlyInstructionsWithLabels)
	check(err)
	parsedAssemblyInstructions := parser.ParseAssemblyInstructions(assemblyInstructionsWithoutLabels)

	fmt.Println(parsedAssemblyInstructions)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
