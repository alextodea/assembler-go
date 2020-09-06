package data

// DestInstructTable is used for decoding of destination symbols within C instructions
var DestInstructTable = map[string]int{
	"null": 0b000,
	"M":    0b001,
	"D":    0b010,
	"MD":   0b011,
	"A":    0b100,
	"AM":   0b101,
	"AD":   0b110,
	"AMD":  0b111,
}

// JumpInstructTable is used for decoding of jump symbols within C instructions
var JumpInstructTable = map[string]int{
	"null": 0b000,
	"JGT":  0b001,
	"JEQ":  0b010,
	"JGE":  0b011,
	"JLT":  0b100,
	"JNE":  0b101,
	"JLE":  0b110,
	"JMP":  0b111,
}

// CompInstructTable is used for decoding of comp symbols within C instructions
var CompInstructTable = map[string]int{
	"0":   0b0101010,
	"1":   0b0111111,
	"-1":  0b0111010,
	"D":   0b0001100,
	"A":   0b0110000,
	"M":   0b1110000,
	"!D":  0b0001101,
	"!A":  0b0110001,
	"!M":  0b1110001,
	"-D":  0b0001111,
	"-A":  0b0110011,
	"-M":  0b1110011,
	"D+1": 0b0011111,
	"A+1": 0b0110111,
	"M+1": 0b1110111,
	"D-1": 0b0001110,
	"A-1": 0b0110010,
	"M-1": 0b1110010,
	"D+A": 0b0000010,
	"D+M": 0b1000010,
	"D-A": 0b0010011,
	"D-M": 0b1010011,
	"A-D": 0b0000111,
	"M-D": 0b1000111,
	"D&A": 0b0000000,
	"D&M": 0b1000000,
	"D|A": 0b0010101,
	"D|M": 0b1010101,
}
