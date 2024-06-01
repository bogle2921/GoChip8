package main

import (
	c "GoChip8/src/include"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		c.CHIP8_WIN_TITLE,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		c.CHIP8_WIDTH,
		c.CHIP8_HEIGHT,
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
	rect := sdl.Rect{0, 0, 100, 100}
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
				break
			}
		}
	}

}
