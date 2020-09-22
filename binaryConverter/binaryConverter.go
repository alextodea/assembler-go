package code

import (
	"assembler-go/data"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ConvertAssemblyToBinary converts assembly instructions to binary
func ConvertAssemblyToBinary(binaryInstructions []uint16, outputF *os.File) {
	for _, binaryValue := range binaryInstructions {
		stringValue := fmt.Sprintf("%016b\n", binaryValue<<0)
		outputF.WriteString(stringValue)
	}
	return
}

// TranslateAssemblyInstructionToBinary translates an assembly instruction to binary code
func TranslateAssemblyInstructionToBinary(parserOutcome map[string]string) (uint16, error) {

	instructionType := parserOutcome["instructionType"]
	var binaryInstruction uint16

	if instructionType == "A" {
		stringSymblolToInteger, err := strconv.Atoi(parserOutcome["symbol"])

		if err != nil {
			return uint16(0), errors.New("symbol string could not be converted to integer")
		}

		binaryInstruction = uint16(stringSymblolToInteger)
		return binaryInstruction, nil

	} else if instructionType == "C" {
		instructionMarker := uint16(0b111 << 13) /// 1110000000000000
		parsedDest := parserOutcome["dest"]
		parsedComp := parserOutcome["comp"]
		parsedJmp := parserOutcome["jump"]

		dest := data.DestInstructTable[parsedDest]
		comp := data.CompInstructTable[parsedComp]
		jmp := data.JumpInstructTable[parsedJmp]

		positionedComp := uint16(comp << 6)
		positionedDest := uint16(dest << 3)
		positionedJmp := uint16(jmp << 0)
		instruction := instructionMarker | positionedComp | positionedDest | positionedJmp

		return instruction, nil
	}
	// else if instructionType == "L" {
	// 	test := parserOutcome["symbol"]
	// }

	return uint16(0), errors.New("instruction type cannot be identified")
}
