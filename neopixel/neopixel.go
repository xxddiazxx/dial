package neopixel

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

var Off = color.RGBA{R: 0x00, G: 0x00, B: 0x00}

type Neopixel struct {
	numLEDs int
	device  ws2812.Device
}

func New(pin machine.Pin, ledCount int) *Neopixel {
	w := ws2812.New(pin)
	w.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	var pixels []color.RGBA
	for range ledCount {
		pixels = append(pixels, Off)
	}

	return &Neopixel{numLEDs: ledCount, device: w}
}

func (n *Neopixel) WriteColors(buf []color.RGBA) (err error) {
	return n.device.WriteColors(buf)
}
