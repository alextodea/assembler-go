package parser

import (
	symbolsHandler "assembler-go/symbolsHandler"
	"strconv"
	"strings"
)

// ParsedInstruction is a type for struct containing instruction type and value
type ParsedInstruction struct {
	instructionType string
	value           []string
}

// ParseAssemblyInstructions reads an assembly language command, parses it, and provides convenient access to the command’s components (fields and symbols). In addition, removes all white space and comments.
func ParseAssemblyInstructions(assemblyInstructionsWithoutLabels []string) []ParsedInstruction {

	parsedAssemblyInstructions := []ParsedInstruction{}

	for _, instruction := range assemblyInstructionsWithoutLabels {
		isAInstruction := strings.Contains(instruction, "@")

		if isAInstruction {
			decodedInstruction := decodeAInstruction(instruction)
			aInstruction := ParsedInstruction{instructionType: "A", value: []string{decodedInstruction}}
			parsedAssemblyInstructions = append(parsedAssemblyInstructions, aInstruction)
		} else {
			dest, comp, jump := decodeCInstruction(instruction)
			cInstruction := ParsedInstruction{instructionType: "C", value: []string{comp, dest, jump}}
			parsedAssemblyInstructions = append(parsedAssemblyInstructions, cInstruction)
		}
	}

	return parsedAssemblyInstructions
}

func decodeCInstruction(textInput string) (dest, comp, jump string) {
	equalSignIndex := strings.Index(textInput, "=")
	semicolonSignIndex := strings.Index(textInput, ";")

	var splittedCInstruction []string

	if equalSignIndex > -1 {
		splittedCInstruction = strings.Split(textInput, "=")
		dest = splittedCInstruction[0]
		comp = splittedCInstruction[1]
		jump = "null"
	} else if semicolonSignIndex > -1 {
		splittedCInstruction = strings.Split(textInput, ";")
		jump = splittedCInstruction[1]
		dest = "null"
		comp = splittedCInstruction[0]
	}

	return dest, comp, jump
}

func decodeAInstruction(aInstruction string) (decodeAInstruction string) {
	decodedAInstruction := strings.ReplaceAll(aInstruction, "@", "")

	if _, instructionIsVariable := strconv.Atoi(decodedAInstruction); instructionIsVariable == nil {
		return decodedAInstruction
	}

	decodedSymbolOfAInstruction := symbolsHandler.VariableSymbols[decodedAInstruction]
	return strconv.Itoa(decodedSymbolOfAInstruction)

}
