package binary

import (
	parser "assembler-go/parser"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConvertAssemblyToBinary converts parsed assembly instructions to binary
func ConvertAssemblyToBinary(parsedAssemblyInstructions [][]string, fileName string) (err error) {

	var instructionToBinary uint16

	outputFileName := strings.Split(fileName, ".")[0]
	outputFileNameAndExtension := outputFileName + ".hack"
	// directoryPath := "../data/binary/"
	// joinFileAndDirPaths := filepath.Join(directoryPath, outputFileNameAndExtension)
	outputFile, err := os.Create(outputFileNameAndExtension)

	if err != nil {
		outputFile.Close()
		return err
	}

	for _, instructionArr := range parsedAssemblyInstructions {
		instructionType := instructionArr[0]

		if instructionType == "A" {
			instructionToInteger, _ := strconv.Atoi(instructionArr[1])
			instructionToBinary = uint16(instructionToInteger)
		} else if instructionType == "C" {
			parsedComp := instructionArr[1]
			parsedDest := instructionArr[2]
			parsedJmp := instructionArr[3]

			comp := parser.CompInstructTable[parsedComp]
			dest := parser.DestInstructTable[parsedDest]
			jmp := parser.JumpInstructTable[parsedJmp]

			// shift bits
			instructionMarker := uint16(0b111 << 13) /// 1110000000000000
			positionedComp := uint16(comp << 6)
			positionedDest := uint16(dest << 3)
			positionedJmp := uint16(jmp << 0)
			instructionToBinary = instructionMarker | positionedComp | positionedDest | positionedJmp // OR
		}

		binaryAsString := fmt.Sprintf("%016b\n", instructionToBinary<<0)
		outputFile.WriteString(binaryAsString)
	}
	outputFile.Close()
	return nil
}

func writeBinaryToFile() {
	outputFileName := strings.Split(os.Args[1], ".")[0]
	outputFileNameAndExtension := outputFileName + ".hack"
	outputF, err := os.Create(outputFileNameAndExtension)

	if err != nil {
		fmt.Println("Failed to convert  assembly lines into binary")
		outputF.Close()
	}

	defer outputF.Close()
}
