// Package pipelines handles necessary data pre-processing through data pipelines
package pipelines

import (
	"bufio"
	"errors"
	"os"
	"path"
	"regexp"
	"strings"
)

// PreProcessData cleans all file input lines and returns an array of assembly commands, including instruction labels
func PreProcessData(fileName string) (assebmlyInstructionsWithLabels []string, err error) {

	filePath := path.Join("./data/assembly/", fileName)
	cleanedFileLines := []string{}

	_, err = os.Stat(filePath)

	if err != nil {
		return cleanedFileLines, err
	}

	f, err := os.Open(filePath)

	if err != nil {
		return cleanedFileLines, err
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		fileLine := input.Text()
		cleanedFileLine, err := cleanFileLine(fileLine)

		if err != nil {
			continue
		}

		cleanedFileLines = append(cleanedFileLines, cleanedFileLine)
	}

	return cleanedFileLines, nil
}

func cleanFileLine(fileLine string) (cleanedFileLine string, err error) {
	fileLinewhiteSpace := regexp.MustCompile(`\s+`)
	fileLineNoWhiteSpace := fileLinewhiteSpace.ReplaceAllString(fileLine, "")
	var isFileLineNoWhiteSpaceCommented bool = strings.HasPrefix(fileLineNoWhiteSpace, "//")

	// if file line is empty or completely commented
	if len(fileLineNoWhiteSpace) == 0 || isFileLineNoWhiteSpaceCommented {
		return "", errors.New("file line empty or completely commented")
	}

	indexOfCommentCharacters := strings.Index(fileLineNoWhiteSpace, "//")

	// if file line contains comment after command
	// then return file line without comment
	if indexOfCommentCharacters > -1 {
		return strings.Split(fileLineNoWhiteSpace, "//")[0], nil
	}

	return fileLineNoWhiteSpace, nil
}

func removeCommentFromAssemblyCommand(textInput string) string {
	i := strings.Index(textInput, "//")

	if i > -1 {
		textInput = strings.Split(textInput, "//")[0]
	}
	return textInput
}

func processInputFileLine(textInput string) (string, error) {
	space := regexp.MustCompile(`\s+`)
	processedTextInput := space.ReplaceAllString(textInput, "")
	var isLineCommented bool = strings.HasPrefix(processedTextInput, "//")

	if len(processedTextInput) == 0 || isLineCommented {
		return "", nil
	}

	return processedTextInput, nil
}
