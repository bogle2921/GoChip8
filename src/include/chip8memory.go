package include

import "strconv"

type Chip8_Memory struct {
	memory [CHIP8_MEM_SIZE]byte
}

func Chip8_mem_set(memory *Chip8_Memory, index uint16, val byte) {
	is_mem_valid(index)
	memory.memory[index] = val
}

func Chip8_mem_get(memory *Chip8_Memory, index uint16) uint8 {
	is_mem_valid(index)
	return memory.memory[index]
}

func is_mem_valid(index uint16) {
	if !(index < CHIP8_MEM_SIZE) {
		panic(strconv.Itoa(int(index)) + " is out of range")
	}
}
