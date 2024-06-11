package include

type Chip8 struct {
	Memory    Memory
	Registers Registers
	Stack     Stack
	Keyboard  Keyboard
}

func Chip8_init(chip8 *Chip8){
	for i := 0; i < int(CHIP8_MEM_SIZE); i++ {
		Mem_set(&chip8.Memory, uint16(i), 0)
	}
}
