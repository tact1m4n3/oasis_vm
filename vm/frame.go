package vm

import "oasis_vm/object"

type Frame struct {
	stack [StackSize]object.Object

	locals []object.Object

	sp int
}

func NewFrame() *Frame {
	frame := &Frame{}
	frame.sp = StackSize - 1

	return frame
}
