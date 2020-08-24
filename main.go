package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	countFileLines(f)
	f.Close()
}

func countFileLines(f *os.File) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		var inputTxt string = input.Text()

		inputTxt = removeLineWhiteSpace(inputTxt)

		var isStringEmpty bool = len(inputTxt) == 0
		var isLineCommented bool = strings.Contains(inputTxt, "//")

		if !isStringEmpty && !isLineCommented {
			fmt.Printf("%s\n", inputTxt)
		}
	}
}

func removeLineWhiteSpace(lineText string) string {
	space := regexp.MustCompile(`\s+`)
	formattedString := space.ReplaceAllString(lineText, "")
	return formattedString
}
