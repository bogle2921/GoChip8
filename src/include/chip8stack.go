package include

type Chip8_Stack struct {
	Stack [CHIP8_STACK_DEPTH]uint16
}

func is_stack_in_bounds(chip8 *Chip8) {
	if !(chip8.Registers.SP < uint16(CHIP8_STACK_DEPTH)) {
		panic("Stack Pointer out of Stack Range")
	}
}

func Chip8_stack_push(chip8 *Chip8, val uint16) {
	is_stack_in_bounds(chip8)
	chip8.Stack.Stack[chip8.Registers.SP] = val
	chip8.Registers.SP += 1
}

func Chip8_stack_pop(chip8 *Chip8) uint16 {
	chip8.Registers.SP -= 1
	is_stack_in_bounds(chip8)
	val := chip8.Stack.Stack[chip8.Registers.SP]
	return val
}
