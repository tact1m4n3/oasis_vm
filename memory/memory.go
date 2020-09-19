package memory

type MemoryInterface interface {
	SetUint8(offset int, value int)
	GetUint8(offset int) byte
	SetUint16(offset int, value int)
	GetUint16(offset int) int
}

func CreateMemory(sizeInBytes int) *Memory {
	arr := make(Memory, sizeInBytes)

	return &arr
}

type Memory []byte

func (memory *Memory) SetUint8(offset int, value int) {
	(*memory)[offset] = byte(value)
}

func (memory *Memory) GetUint8(offset int) byte {
	return (*memory)[offset]
}

func (memory *Memory) SetUint16(offset int, value int) {
	memory.SetUint8(offset, value>>8)
	memory.SetUint8(offset+1, value&0xFF)
}

func (memory *Memory) GetUint16(offset int) int {
	return (int((*memory)[offset]) << 8) | int(((*memory)[offset+1]))
}
