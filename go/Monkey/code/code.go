package code

import (
	"encoding/binary"
	"fmt"
)

type Instrcutions []byte

type Opecode byte

const (
	OpConstant Opecode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opecode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opecode(op)]
	if !ok {
		return nil, fmt.Errorf("opecode %d undefined", op)
	}

	return def, nil
}

func Make(op Opecode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}
