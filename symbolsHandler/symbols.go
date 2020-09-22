// Package symbolshandler handles assembly instruction labels and variabler symbols
package symbolshandler

// VariableSymbols is used to store variable names and pointers to their locations in memory
var VariableSymbols = map[string]int{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16386,
	"KBD":    24576,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
}

// MemoryCounter keeps track of next available word in memory
var MemoryCounter int = 16

// LabelSymbols is used in order to store instruction memory locations that can be used as references
var LabelSymbols = map[string]int {}