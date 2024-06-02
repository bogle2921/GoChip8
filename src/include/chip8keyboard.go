package include

type Chip8_Keyboard struct {
	Keyboard [CHIP8_NUM_KEYS]bool
}

func is_key_valid(key byte) {
	if !(key < CHIP8_NUM_KEYS) {
		panic("Key pressed is out of range")
	}
}

func Chip8_keyboard_map(key_map []int, key byte) int {
	for i := 0; i < int(CHIP8_NUM_KEYS); i++ {
		if key_map[i] == int(key) {
			return i
		}
	}
	return -1
}

func Chip8_key_press(keyboard *Chip8_Keyboard, key byte) {
	is_key_valid(key)
	keyboard.Keyboard[key] = true
}

func Chip8_key_up(keyboard *Chip8_Keyboard, key byte) {
	is_key_valid(key)
	keyboard.Keyboard[key] = false
}

func Chip8_is_key_pressed(keyboard *Chip8_Keyboard, key byte) bool {
	is_key_valid(key)
	return keyboard.Keyboard[key]
}
