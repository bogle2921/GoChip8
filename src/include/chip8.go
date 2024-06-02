package include

type Chip8 struct {
	Memory    Chip8_Memory
	Registers Chip8_Registers
	Stack     Chip8_Stack
	Keyboard Chip8_Keyboard
}
