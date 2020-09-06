package main

import (
	binaryConverter "assembler-go/binaryConverter"
	parser "assembler-go/parser"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	binaryInstructions, err := parseFileLines(f)

	if err != nil {
		f.Close()
	}

	outputFileName := strings.Split(fileName, ".")[0]
	outputFileNameAndExtension := outputFileName + ".hack"
	outputF, err := os.Create(outputFileNameAndExtension)

	if err != nil {
		fmt.Println("Failed to convert  assembly lines into binary")
		outputF.Close()
	}

	defer f.Close()

	convertAssemblyToBinary(binaryInstructions, outputF)
	fmt.Println("Succesfully generated binary commands and saved them into " + outputFileNameAndExtension)
}

func convertAssemblyToBinary(binaryInstructions []uint16, outputF *os.File) {
	for _, binaryValue := range binaryInstructions {
		stringValue := fmt.Sprintf("%016b\n", binaryValue<<0)
		outputF.WriteString(stringValue)
	}
	return
}

func parseFileLines(f *os.File) ([]uint16, error) {
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

		binaryInstructions = append(binaryInstructions, binaryInstruction)
	}

	if len(binaryInstructions) < 1 {
		return binaryInstructions, errors.New("binary instructions file is empty")
	}

	return binaryInstructions, nil
}
