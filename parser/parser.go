package parser

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	commandA string = "A_COMMAND"
	commandL string = "L_COMMAND"
	commandC string = "C_COMMAND"
)

// ParseInstruction reads an assembly language command, parses it, and provides convenient access to the command’s components (fields and symbols). In addition, removes all white space and comments.
func ParseInstruction(assemblyCommand string) {
	var formattedAssemblyCommand string = removeLineWhiteSpace(assemblyCommand)

	var isStringEmpty bool = len(formattedAssemblyCommand) == 0
	var isLineCommented bool = strings.Contains(formattedAssemblyCommand, "//")

	if !isStringEmpty && !isLineCommented {
		var commandType string = determineCommandType(assemblyCommand)
		fmt.Printf(commandType)

	}
}

func removeLineWhiteSpace(textInput string) string {
	space := regexp.MustCompile(`\s+`)
	formattedString := space.ReplaceAllString(textInput, "")
	return formattedString
}

func determineCommandType(assemblyCommand string) string {
	var isCommandOfTypeA bool = isCommandOfTypeA(assemblyCommand)
	var isCommandOfTypeL bool = isCommandOfTypeL(assemblyCommand)

	switch true {
	case isCommandOfTypeA:
		return commandA
	case isCommandOfTypeL:
		return commandL
	default:
		return commandC
	}
}

func isCommandOfTypeA(assemblyCommand string) bool {

	i := strings.Index(assemblyCommand, "@")

	if i == 0 {
		return true
	}
	return false
}

func isCommandOfTypeL(assemblyCommand string) bool {
	i := strings.Index(assemblyCommand, "(")

	if i == 0 {
		return true
	}
	return false
}
