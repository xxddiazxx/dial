package main

import (
	"fmt"
	"image/color"
	"machine"
	"machine/usb/hid/mouse"

	"github.com/xxddiazxx/dial/encoder"
	"github.com/xxddiazxx/dial/neopixel"
)

func main() {
	fmt.Println("Initializing")

	n := neopixel.New(machine.A1, 17) // I didnt pay attention and cut only 17 LEDS, but you should have 18 if you followed the instructions. Don't be me.

	yellow := color.RGBA{R: 0xff, G: 0xff, B: 0x00}
	allYellow := []color.RGBA{yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow, yellow} // 17 LEDS
	n.WriteColors(allYellow)

	e := encoder.New()

	lastPosition := 0
	lastValue := encoder.ButtonReleased

	// main loop
	for {
		position, err := e.Position()
		if err != nil {
			fmt.Println("error:", err)
		}

		value, err := e.Button.Value()
		if err != nil {
			fmt.Println("error:", err)
		}

		if position != lastPosition {
			if position > lastPosition {
				fmt.Println("greater", position)
				mouse.Mouse.WheelUp()
			} else {
				fmt.Println("less", position)
				mouse.Mouse.WheelDown()
			}

			lastPosition = position
		}

		if value != lastValue {
			if value == encoder.ButtonPressed {
				fmt.Println("pressed")
			} else {
				fmt.Println("released")
			}
			lastValue = value
		}

		// w.WriteColors(colorOn)
		// time.Sleep(time.Millisecond * 1000)

		// w.WriteColors(colorOff)
		// time.Sleep(time.Millisecond * 1000)
	}
}

// TODO ==================================================================================================
// ✔ Get LED strip to work
// ✔ Get Mouse scroll keys to work
// ✔ Get Encoder to report position correctly (mirror python rollback)
// ✔ Get Button on encoder to register on_pressed
// ☐ Get Keyboard keys to work
// ☐ Get Debounce, hold and tap logic working (reduce time to read?) Note: I think adafruit handled debounce with the module. Thanks Adafriut team!
// ☐ Attach a buzzer: https://pkg.go.dev/tinygo.org/x/drivers/buzzer (I want it to click when I rotate it)
// ✔ Cleanup
// ☐ Disable mounting serial

// Nice to have:
// set position?
// Get LED on Stemma working
// Get LED on main mc working
// Clean up getPosition to limit the number of conversions

// Usefull links:
// https://tinygo.org/docs/reference/microcontrollers/qtpy-rp2040/
// https://github.com/adafruit/Adafruit_Seesaw
// https://tinygo.org/docs/concepts/peripherals/i2c/
// https://pkg.go.dev/tinygo.org/x/drivers@v0.30.0/seesaw
// https://adafruit.github.io/Adafruit_Seesaw/html/class_adafruit__seesaw.html
// https://www.youtube.com/@pragmatiktech
// https://github.com/tinygo-org/drivers/blob/v0.30.0/examples/seesaw/main.go <-- Seesaw example
// https://github.com/adafruit/Adafruit_CircuitPython_seesaw/blob/main/adafruit_seesaw/digitalio.py

// Still Need:
// https://learn.adafruit.com/usb-rotary-media-dial/code-the-media-dial
// https://github.com/adafruit/Adafruit_CircuitPython_Debouncer/blob/main/adafruit_debouncer.py#L124
// https://github.com/adafruit/Adafruit_CircuitPython_Debouncer/blob/5932ce49e5a4fef6f4c3a5d17d0d392febc937f7/adafruit_debouncer.py#L4

// Python Playground
// https://www.online-python.com/KXmLvup8J6

// Go Playground
// https://go.dev/play/p/X-nPOi7lUHY
