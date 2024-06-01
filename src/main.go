package main

import (
	c "GoChip8/src/include"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func main() {
	chip8 := c.Chip8{}
	chip8.Registers.SP = 0
	c.Chip8_stack_push(&chip8, 0xff)
	c.Chip8_stack_push(&chip8, 0xaa)
	println("%d", c.Chip8_stack_pop(&chip8))
	println("%d", c.Chip8_stack_pop(&chip8))
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
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}

}
