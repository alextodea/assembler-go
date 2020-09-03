package main

import (
	binaryConverter "assembler-go/binaryConverter"
	parser "assembler-go/parser"
	"bufio"
	"errors"
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

	binaryInstructions, err := countFileLines(f)

	if err != nil {
		f.Close()
	}

	outputF, err := os.Create(fileName + ".hack")

	if err != nil {
		outputF.Close()
	}

	defer f.Close()

	for binaryInstruction := range binaryInstructions {
		outputF.WriteString(fmt.Sprintf("%15b\n", binaryInstruction))
	}
}

func countFileLines(f *os.File) ([]uint16, error) {
	input := bufio.NewScanner(f)
	var binaryInstructions []uint16

	for input.Scan() {
		fileLine := input.Text()
		parserOutcome, err := parser.ParseInstruction(fileLine)

		if err != nil {
			fmt.Println(err)
			continue
		}

		binaryInstruction, err := binaryConverter.TranslateAssemblyInstructionToBinary(parserOutcome)

		if err != nil {
			f.Close()
		}

		if binaryInstruction > 1 {
			binaryInstructions = append(binaryInstructions, binaryInstruction)
		}
	}

	if len(binaryInstructions) < 1 {
		return binaryInstructions, errors.New("binary instructions file is empty")
	}

	return binaryInstructions, nil
}
