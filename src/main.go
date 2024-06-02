package main

import (
	c "GoChip8/src/include"

	sdl "github.com/veandco/go-sdl2/sdl"
)

var keyboard_map = [c.CHIP8_NUM_KEYS]int{
	sdl.K_0, sdl.K_1, sdl.K_2, sdl.K_3, sdl.K_4, sdl.K_5,
	sdl.K_6, sdl.K_7, sdl.K_8, sdl.K_9, sdl.K_a, sdl.K_b,
	sdl.K_c, sdl.K_d, sdl.K_e, sdl.K_f,
}

func main() {
	chip8 := c.Chip8{}

	c.Chip8_key_press(&chip8.Keyboard, 0x0f)
	c.Chip8_key_up(&chip8.Keyboard, 0x0f)
	is_down := c.Chip8_is_key_pressed(&chip8.Keyboard, 0x0f)
	println(is_down)

	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		c.CHIP8_WIN_TITLE,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		c.CHIP8_WIDTH*c.CHIP8_WIN_MULTI,
		c.CHIP8_HEIGHT*c.CHIP8_WIN_MULTI,
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	surface, err := win.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0)
	rect := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}
	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	surface.FillRect(&rect, pixel)
	win.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			case *sdl.KeyboardEvent:
				keystate := e.Type
				switch keystate {
				case sdl.KEYUP:
					println("Key up")
				case sdl.KEYDOWN:
					println("Key down")
				}
			}
		}
	}

}
