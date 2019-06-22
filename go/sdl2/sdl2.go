package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type color struct {
	r, g, b byte
}

const winWidth int = 800
const winHeight int = 600

func main() {
	fmt.Println("start")
	window, err := sdl.CreateWindow("SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	fmt.Println("window created")
	ren, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer ren.Destroy()

	fmt.Println("renderer created")
	tex, err := ren.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING,
		int32(winWidth), int32(winHeight))
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()
	fmt.Println("tex created")
	pixels := make([]byte, winWidth*winHeight*4)
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{255, 0, 255}, pixels)
		}
	}

	tex.Update(nil, pixels, winWidth*4)
	ren.Copy(tex, nil, nil)
	ren.Present()
	sdl.Delay(2000)
}

func setPixel(x, y int, c color, pixels []byte) {
	fmt.Println("set pixel called", x, y)
	index := (y*winWidth + x) * 4
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}
