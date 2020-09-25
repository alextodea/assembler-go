// Package symbolshandler handles assembly instruction labels and variabler symbols
package symbolshandler

import (
	"regexp"
	"strings"
)

// HandleSymbols maps labels to their command instruction line numbers, removes labels from instruction set and also maps memory locations to variables
func HandleSymbols(assebmlyInstructionsWithLabels []string) (assebmlyInstructions []string, err error) {
	assemblyInstructionsNoLabels := mapLabelsToSymbolsTable(assebmlyInstructionsWithLabels)
	mapAddressesToSymbolTable(assemblyInstructionsNoLabels)

	return assemblyInstructionsNoLabels, nil
}

func mapLabelsToSymbolsTable(assebmlyInstructionsWithLabels []string) []string {
	assemblyInstructionsNoLabels := []string{}
	labelsCounter := 0
	for index, instruction := range assebmlyInstructionsWithLabels {
		indexOfLabelOpenParantheses := strings.Index(instruction, "(")
		if indexOfLabelOpenParantheses > -1 {
			label := strings.ReplaceAll(instruction, "(", "")
			label = strings.ReplaceAll(label, ")", "")
			SymbolsTable[label] = index - labelsCounter
			labelsCounter++
			continue
		}

		assemblyInstructionsNoLabels = append(assemblyInstructionsNoLabels, instruction)
	}
	return assemblyInstructionsNoLabels
}

func mapAddressesToSymbolTable(assemblyInstructionsNoLabels []string) {
	for _, instruction := range assemblyInstructionsNoLabels {
		isInstructionOfTypeA := strings.Index(instruction, "@")
		if isInstructionOfTypeA < 0 {
			continue
		}

		address := strings.Split(instruction, "@")[1]
		isSymbol := isLetter(address[0:1])

		if !isSymbol {
			continue
		}

		if _, addressExistsInMemory := SymbolsTable[address]; addressExistsInMemory {
			continue
		}

		SymbolsTable[address] = MemoryCounter
		MemoryCounter++
		continue
	}
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
