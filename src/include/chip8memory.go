package include

type Chip8_Memory struct {
	memory [CHIP8_MEM_SIZE]byte
}

func Chip8_mem_set(memory *Chip8_Memory, index uint16, val byte) {
	memory.memory[index] = val
}

func Chip8_mem_get(memory *Chip8_Memory, index uint16) uint8 {
	return memory.memory[index]
}

/*func is_mem_valid(index uint16) {
	if !(index >= 0 && index < CHIP8_MEM_SIZE) {
		panic(strconv.Itoa(int(index)) + " is out of range")
	}
} Not necessary due to built in uint16 check */
