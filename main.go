package main

import (
	pipelines "assembler-go/pipelines"
	symbolshandler "assembler-go/symbolsHandler"
	"fmt"
	"os"
)

func main() {

	assebmlyInstructionsWithLabels, err := pipelines.PreProcessData(os.Args[1])
	check(err)
	assebmlyInstructionsWithoutLabels, err := symbolshandler.HandleSymbols(assebmlyInstructionsWithLabels)
	check(err)
	fmt.Println(assebmlyInstructionsWithoutLabels)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
