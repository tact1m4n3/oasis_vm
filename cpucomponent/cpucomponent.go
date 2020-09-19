package cpucomponent

import (
	"bytes"
	"oasis_vm/instructions"
	"oasis_vm/memorymapper"
)

const ipAddress int = 0xdff0
const spAddress int = 0xdff2

type CPU struct {
	Memory *memorymapper.MemoryMapper
	ip     int
	sp     int
}

func NewCPU(memory *memorymapper.MemoryMapper) *CPU {
	cpu := &CPU{Memory: memory}
	cpu.ip = 0
	cpu.sp = 0xffff - 1
	cpu.Memory.SetUint16(spAddress, cpu.sp)
	return cpu
}

func (cpu *CPU) ShowMemory(offset int) string {
	var out bytes.Buffer
	for i := 0; i < 8; i++ {
		out.WriteString(string(cpu.Memory.GetUint8(offset + i)))
	}

	return out.String()
}

func (cpu *CPU) fetch() byte {
	instruction := cpu.Memory.GetUint8(cpu.ip)
	cpu.ip++
	cpu.Memory.SetUint16(ipAddress, cpu.ip)
	return instruction
}

func (cpu *CPU) fetch16() int {
	instruction := cpu.Memory.GetUint16(cpu.ip)
	cpu.ip += 2
	cpu.Memory.SetUint16(ipAddress, cpu.ip)
	return instruction
}

func (cpu *CPU) pushInt(value int) {
	cpu.Memory.SetUint16(cpu.sp, value)
	cpu.sp -= 2
	cpu.Memory.SetUint16(spAddress, cpu.sp)
}

func (cpu *CPU) popInt() int {
	cpu.sp += 2
	value := cpu.Memory.GetUint16(cpu.sp)

	return value
}

func (cpu *CPU) popList() []int {
	var list []int
	length := cpu.popInt()

	for i := 1; i <= length; i++ {
		list = append(list, cpu.popInt())
	}

	return list
}

func (cpu *CPU) execute(instruction byte) {
	switch instruction {
	case instructions.PUSH:
		cpu.pushInt(cpu.fetch16())

	/*
		case instructions.PUSH_LIST:
			length := cpu.fetch16()

			for i := 1; i <= length; i++ {
				cpu.pushInt(cpu.fetch16())
			}
			cpu.pushInt(length)
	*/
	case instructions.POP:
		cpu.popInt()

	/*
		case instructions.POP_LIST:
			length := cpu.popInt()

			for i := 1; i <= length; i++ {
				cpu.popInt()
			}
	*/

	case instructions.STORE_MEM:
		addr := cpu.fetch16()
		value := cpu.popInt()

		cpu.Memory.SetUint16(addr, value)

	case instructions.LOAD_MEM:
		addr := cpu.fetch16()
		value := cpu.Memory.GetUint16(addr)

		cpu.pushInt(value)

	case instructions.ADD:
		nr2 := cpu.popInt()
		nr1 := cpu.popInt()

		cpu.pushInt(nr1 + nr2)

	case instructions.SUB:
		nr2 := cpu.popInt()
		nr1 := cpu.popInt()

		cpu.pushInt(nr1 - nr2)

	case instructions.MUL:
		nr2 := cpu.popInt()
		nr1 := cpu.popInt()

		cpu.pushInt(nr1 * nr2)

	case instructions.DIV:
		nr2 := cpu.popInt()
		nr1 := cpu.popInt()

		cpu.pushInt(nr1 / nr2)
	}
}

// Run executes the instructions
func (cpu *CPU) Run() {
	for {
		instruction := cpu.fetch()
		cpu.execute(instruction)
		if instruction == instructions.HLT {
			break
		}
	}
}
