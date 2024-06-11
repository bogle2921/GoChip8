package include

type Registers struct {
	V           [CHIP8_NUM_REGISTERS]byte
	I           uint16
	Delay_timer byte
	Sound_timer byte
	PC          uint16
	SP          uint16
}
