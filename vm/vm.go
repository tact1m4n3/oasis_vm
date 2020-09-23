package vm

import (
	"oasis_vm/code"
	"oasis_vm/object"
)

const StackSize = 2048
const MaxFrames = 64

type VM struct {
	constants []object.Object
	globals   []object.Object

	instructions code.Instructions

	frames      [MaxFrames]*Frame
	framesIndex int

	ip int
}

func NewVM(instructions code.Instructions) *VM {
	vm := &VM{instructions: instructions}
	vm.framesIndex = MaxFrames - 1

	mainFrame := NewFrame()
	mainFrame.sp = StackSize - 1

	vm.pushFrame(mainFrame)

	return vm
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framesIndex]
}

func (vm *VM) pushFrame(frame *Frame) {
	vm.frames[vm.framesIndex] = frame
	vm.framesIndex--
}

func (vm *VM) popFrame() {
	vm.framesIndex--
}

func (vm *VM) push(obj object.Object) {
	vm.currentFrame().stack[vm.currentFrame().sp] = obj
	vm.currentFrame().sp--
}

func (vm *VM) pop() object.Object {
	vm.currentFrame().sp++
	return vm.currentFrame().stack[vm.currentFrame().sp]
}

func (vm *VM) fetchInstruction() *code.Definition {
	bytecode := vm.instructions[vm.ip]

	definition, err := code.Lookup(bytecode)
	if err != nil {
		panic(err)
	}

	vm.ip++

	return definition
}

func (vm *VM) executeInstruction(instruction *code.Definition) {
	switch instruction.Name {
	case "IntConstant":
		value := code.GetInt64(vm.instructions[vm.ip:])

		integer := &object.Integer{Value: value}
		vm.constants = append(vm.constants, integer)

	case "OpConstant":
		constIdx := code.GetUint16(vm.instructions[vm.ip:])

		vm.push(vm.constants[constIdx])
	}
}

func (vm *VM) run() {
	for i := 0; i < len(vm.instructions); i++ {
		instruction := vm.fetchInstruction()
		vm.executeInstruction(instruction)
	}
}
