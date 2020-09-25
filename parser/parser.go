package parser

import (
	symbolsHandler "assembler-go/symbolsHandler"
	"strconv"
	"strings"
)

// ParseAssemblyInstructions reads an assembly language command, parses it, and provides convenient access to the command’s components (fields and symbols). In addition, removes all white space and comments.
func ParseAssemblyInstructions(assemblyInstructionsWithoutLabels []string) [][]string {

	parsedAssemblyInstructions := [][]string{}

	for _, instruction := range assemblyInstructionsWithoutLabels {
		isAInstruction := strings.Contains(instruction, "@")

		if isAInstruction {
			decodedInstruction := decodeAInstruction(instruction)
			parsedAssemblyInstructions = append(parsedAssemblyInstructions, []string{"A", decodedInstruction})
		} else {
			dest, comp, jump := decodeCInstruction(instruction)
			parsedAssemblyInstructions = append(parsedAssemblyInstructions, []string{"C", comp, dest, jump})
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

	decodedSymbolOfAInstruction := symbolsHandler.SymbolsTable[decodedAInstruction]
	return strconv.Itoa(decodedSymbolOfAInstruction)

}
