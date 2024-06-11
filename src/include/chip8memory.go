package include

import "strconv"

type Memory struct {
	memory [CHIP8_MEM_SIZE]byte
}

func Mem_set(memory *Memory, index uint16, val byte) {
	is_mem_valid(index)
	memory.memory[index] = val
}

func Mem_get(memory *Memory, index uint16) uint8 {
	is_mem_valid(index)
	return memory.memory[index]
}

func is_mem_valid(index uint16) {
	if !(index < CHIP8_MEM_SIZE) {
		panic(strconv.Itoa(int(index)) + " is out of range")
	}
}
