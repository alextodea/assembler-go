package code

import (
	"errors"
	"fmt"
	"strconv"
)

// TranslateAssemblyInstructionToBinary translates an assembly instruction to binary code
func TranslateAssemblyInstructionToBinary(parserOutcome map[string]string) (uint16, error) {

	instructionType := parserOutcome["instructionType"]
	var binaryInstruction uint16

	fmt.Println(binaryInstruction)
	if instructionType == "A" {
		stringSymblolToInteger, err := strconv.Atoi(parserOutcome["symbol"])

		if err != nil {
			return uint16(0), errors.New("symbol string could not be converted to integer")
		}

		binaryInstruction = uint16(stringSymblolToInteger)
	}

	return binaryInstruction, nil

	//  convert binary instruction to string and add all zeros
	// i.e. stringBI = fmt.printf("%15b", binaryInstruction)
	// otherwise filter the incoming lines to be only the right ones
}
