package parser

import (
	"errors"
	"regexp"
	"strings"
)

// ParseInstruction reads an assembly language command, parses it, and provides convenient access to the command’s components (fields and symbols). In addition, removes all white space and comments.
func ParseInstruction(assemblyCommand string) (map[string]string, error) {
	assemblyCommand, err := processInputFileLine(assemblyCommand)

	if err != nil {
		return nil, err
	}

	assemblyCommand = removeCommentFromAssemblyCommand(assemblyCommand)

	instructionType, err := determineInstructionType(assemblyCommand)

	if err != nil {
		return nil, err
	}

	parserOutcome := map[string]string{}
	parserOutcome["instructionType"] = instructionType

	switch instructionType {
	case "A":
		parserOutcome["symbol"] = decodeSymbol(assemblyCommand, "@")
		return parserOutcome, nil
	case "C":
		dest, comp, jump := decodeCInstruction(assemblyCommand)
		parserOutcome["dest"] = dest
		parserOutcome["comp"] = comp
		parserOutcome["jump"] = jump
		return parserOutcome, nil
	case "L":
		firstSplit := decodeSymbol(assemblyCommand, "(")
		parserOutcome["symbol"] = decodeSymbol(firstSplit, ")")
		return parserOutcome, nil
	default:
		return map[string]string{}, errors.New("instruction could not be parsed " + assemblyCommand)
	}
}

func determineInstructionType(assemblyCommand string) (string, error) {
	isAInstruction := strings.Contains(assemblyCommand, "@")
	isLInstruction := strings.Contains(assemblyCommand, "(")
	isCInstruction := strings.Contains(assemblyCommand, "=") || strings.Contains(assemblyCommand, ";")

	switch true {
	case isAInstruction && !isCInstruction && !isLInstruction:
		return "A", nil
	case !isAInstruction && isCInstruction && !isLInstruction:
		return "C", nil
	case !isAInstruction && !isCInstruction && isLInstruction:
		return "L", nil
	default:
		return "", errors.New("instruction type cannot be identified")
	}
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
		return "", errors.New("didn't parsed this line: " + processedTextInput)
	}

	return processedTextInput, nil
}

func removeCommentFromAssemblyCommand(textInput string) string {
	i := strings.Index(textInput, "//")

	if i > -1 {
		textInput = strings.Split(textInput, "//")[0]
	}
	return textInput

}
