package code

import (
	dataModule "assembler-go/data"
	"fmt"
)

// TranslateAssemblyInstructionToBinary translates an assembly instruction to binary code
func TranslateAssemblyInstructionToBinary(parsedInstructionComponents map[string]string) (string, error) {

	if len(parsedInstructionComponents) < 2 {
		return "", fmt.Errorf("parsedInstructionComponents does not contain enough elements: %v", parsedInstructionComponents)
	}

	instructionType := parsedInstructionComponents["type"]

	if instructionType == "C" {
		dest := parsedInstructionComponents["dest"]
		comp := parsedInstructionComponents["comp"]
		jump := parsedInstructionComponents["jump"]
		binaryCInstruction := "111" + dataModule.CompInstructTable[comp] + dataModule.DestInstructTable[dest] + dataModule.JumpInstructTable[jump]
		return binaryCInstruction, nil
	}

	if instructionType == "A" {
		symbol := parsedInstructionComponents["symbol"]
		// symbolToInt, _ := strconv.Atoi(symbol)
		fmt.Print(symbol)
		// intSymbolTo64 := int64(symbolToInt)
		// binarySymbol := strconv.FormatInt(intSymbolTo64, 2)
		// fmt.Printf("A: %d: %016b", symbolToInt, symbolToInt)
		binaryAInstruction := "0"
		return binaryAInstruction, nil
	}

	return instructionType, nil
}
