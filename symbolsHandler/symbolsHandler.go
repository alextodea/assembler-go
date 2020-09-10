// Package symbolshandler handles assembly instruction labels and variabler symbols
package symbolshandler

import (
	"fmt"
	"regexp"
	"strings"
)

// HandleSymbols maps labels to their command instruction line numbers, removes labels from instruction set and also maps memory locations to variables
func HandleSymbols(assebmlyInstructionsWithLabels []string) (assebmlyInstructions []string, err error) {
	cleanAssemblyInstructions := []string{}
	indexAfterRemovingLabels := 0
	for _, instruction := range assebmlyInstructionsWithLabels {
		indexOfLabelOpenParantheses := strings.Index(instruction, "(")
		if indexOfLabelOpenParantheses > -1 {
			handleLabel(instruction, indexAfterRemovingLabels)
			continue
		}

		indexOfAddressSymbol := strings.Index(instruction, "@")

		if indexOfAddressSymbol > -1 {
			handleSymbol(instruction, indexAfterRemovingLabels)
		}

		cleanAssemblyInstructions = append(cleanAssemblyInstructions, instruction)
		indexAfterRemovingLabels++
	}
	fmt.Println("LabelSymbols", LabelSymbols)
	fmt.Println("VariableSymbols", VariableSymbols)
	fmt.Println("MemoryCounter", MemoryCounter)
	return cleanAssemblyInstructions, nil
}

func handleLabel(label string, index int) {
	LabelSymbols[label] = index
}

func handleSymbol(instruction string, index int) {
	address := strings.Split(instruction, "@")[1]
	isSymbol := isLetter(address[0:1])

	if !isSymbol {
		return
	}

	if _, addressExistsInMemory := VariableSymbols[address]; addressExistsInMemory {
		return
	}

	VariableSymbols[address] = MemoryCounter
	MemoryCounter++
	return
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
