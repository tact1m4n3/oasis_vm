package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	IntConstant Opcode = iota
	StringConstant

	OpConstant
	Add
	Sub
	Mul
	Div
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	IntConstant:    {"IntConstant", []int{}},
	StringConstant: {"StringConstant", []int{}},

	OpConstant: {"OpConstant", []int{2}},

	Add: {"Add", []int{}},
	Sub: {"Sub", []int{}},
	Mul: {"Mul", []int{}},
	Div: {"Div", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func GetInt64(ins Instructions) int64 {
	return int64(binary.BigEndian.Uint64(ins))
}

func GetUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func GetUint8(ins Instructions) uint8 {
	return uint8(ins[0])
}
