package main

import (
	"fmt"
	"io/ioutil"
	"oasis_vm/cpucomponent"
	"oasis_vm/memory"
	"oasis_vm/memorymapper"
	"os"
	"strconv"
	"strings"
)

func main() {
	MM := memorymapper.NewMemoryMapper()

	globalMemory := memory.CreateMemory(256 * 256)
	MM.MapDevice(globalMemory, 0, 0xffff, true)

	MM.MapDevice(globalMemory, 0x2000, 0x2fff, true)
	MM.MapDevice(globalMemory, 0x3000, 0x4fff, true)

	MM.MapDevice(globalMemory, 0xe0000, 0xffff, true)

	cpu := cpucomponent.NewCPU(MM)

	bin, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	binary_code := strings.Split(string(bin), " ")

	for i, instruction := range binary_code {
		instruction, _ := strconv.Atoi(instruction)
		cpu.Memory.SetUint8(i, instruction)
	}

	cpu.Run()

	fmt.Println(cpu.Memory.GetUint16(0xffff - 1))
}
