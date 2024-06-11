package include

type Keyboard struct {
	Keyboard [CHIP8_NUM_KEYS]bool
}

func is_key_valid(key int) {
	if !(key < CHIP8_NUM_KEYS) {
		panic("Key pressed is out of range")
	}
}

func Keyboard_map(key_map [16]int, key int) int {
	for i := 0; i < int(CHIP8_NUM_KEYS); i++ {
		if key_map[i] == int(key) {
			return i
		}
	}
	return -1
}

func Keyboard_key_press(keyboard *Keyboard, key int) {
	is_key_valid(key)
	keyboard.Keyboard[key] = true
}

func Keyboard_key_up(keyboard *Keyboard, key int) {
	is_key_valid(key)
	keyboard.Keyboard[key] = false
}

func Keyboard_is_key_pressed(keyboard *Keyboard, key int) bool {
	is_key_valid(key)
	return keyboard.Keyboard[key]
}
