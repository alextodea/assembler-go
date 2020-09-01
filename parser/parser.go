package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// receive file line as input
// replace white space with no space
// skip line if commented or empty

// ParseInstruction reads an assembly language command, parses it, and provides convenient access to the command’s components (fields and symbols). In addition, removes all white space and comments.
func ParseInstruction(assemblyCommand string) (map[string]string, error) {
	assemblyCommand, err := processInputFileLine(assemblyCommand)

	if err != nil {
		return nil, err
	}

	fmt.Println(assemblyCommand)

	// var isStringEmpty bool = len(assemblyCommandWithoutWhiteSpace) == 0
	// var isLineCommented bool = strings.HasPrefix(assemblyCommandWithoutWhiteSpace, "//")

	// if !isStringEmpty && !isLineCommented {

	// 	var formattedAssemblyLine string = removeCommentFromAssemblyCommand(assemblyCommandWithoutWhiteSpace)

	// 	isAInstruction := strings.Contains(formattedAssemblyLine, "@")
	// 	isLInstruction := strings.Contains(formattedAssemblyLine, "(")
	// 	isCInstruction := strings.Contains(formattedAssemblyLine, "=") || strings.Contains(formattedAssemblyLine, ";")

	parserOutcome := make(map[string]string)
	// 	switch true {
	// 	case isAInstruction:
	// 		parserOutcome["type"] = "A"
	// 		parserOutcome["symbol"] = decodeSymbol(formattedAssemblyLine, "@")
	// 		return parserOutcome, nil
	// 	case isCInstruction:
	// 		parserOutcome["type"] = "C"
	// 		dest, comp, jump := decodeCInstruction(formattedAssemblyLine)
	// 		parserOutcome["dest"] = dest
	// 		parserOutcome["comp"] = comp
	// 		parserOutcome["jump"] = jump
	// 		return parserOutcome, nil
	// 	case isLInstruction:
	// 		parserOutcome["type"] = "L"
	// 		firstSplit := decodeSymbol(formattedAssemblyLine, "(")
	// 		parserOutcome["symbol"] = decodeSymbol(firstSplit, ")")
	// 		return parserOutcome, nil
	// 	default:
	// 		return map[string]string{}, fmt.Errorf("command could bit be parser: %v", formattedAssemblyLine)
	// 	}
	// }

	return parserOutcome, nil
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

func decodeSymbol(textInput, oldChar string) string {
	return strings.ReplaceAll(textInput, oldChar, "")
}

func processInputFileLine(textInput string) (string, error) {
	space := regexp.MustCompile(`\s+`)
	processedTextInput := space.ReplaceAllString(textInput, "")
	var isLineCommented bool = strings.HasPrefix(processedTextInput, "//")

	if len(processedTextInput) == 0 || isLineCommented {
		return "", errors.New("file line is either commented or empty: " + processedTextInput)
	}

	return processedTextInput, nil
}

func removeCommentFromAssemblyCommand(textInput string) string {
	i := strings.Index(textInput, "//")

	if i > -1 {
		splittedTextInput := strings.Split(textInput, "//")
		return splittedTextInput[0]
	}
	return textInput

}
