package encoder

import (
	"encoding/binary"
	"machine"

	"tinygo.org/x/drivers/seesaw"
)

type Encoder struct {
	device *seesaw.Device
	Button *Button
}

type ButtonState int

const (
	ButtonPressed ButtonState = iota
	ButtonReleased
	ButtonHold
	Unknown
)

type Button struct {
	device       *seesaw.Device
	State        ButtonState
	PressedCount int
}

func New() *Encoder {
	// Configure I2C device
	i2c := machine.I2C1
	i2c.Configure(machine.I2CConfig{
		SCL: machine.I2C1_QT_SCL_PIN,
		SDA: machine.I2C1_QT_SDA_PIN,
	})

	// This is specific to the adafruit encoder
	e := seesaw.New(i2c)
	e.Address = 0x36

	b := seesaw.New(i2c)
	b.Address = 0x36

	return &Encoder{
		device: e,
		Button: &Button{
			device:       b,
			State:        ButtonReleased,
			PressedCount: 0,
		},
	}
}

func (e *Encoder) Position() (int, error) {
	var buf [4]byte
	err := e.device.Read(seesaw.ModuleEncoderBase, 0x30, buf[:])
	if err != nil {
		return 0, err
	}

	// negate the position to make clockwise rotation positive
	return int(int32(binary.BigEndian.Uint32(buf[:])) * -1), nil // TODO: This is ugly, can we fix?
}

func (b *Button) Value() (ButtonState, error) {
	var buf [4]byte
	err := b.device.Read(seesaw.ModuleGpioBase, seesaw.FunctionGpioBulk, buf[:])
	if err != nil {
		return Unknown, err
	}

	pin := (1 << 24) // I am not sure where the example got pin 24, but lets use it
	ret := int(binary.BigEndian.Uint32(buf[:]))

	if (ret & pin) == 0 {
		return ButtonPressed, nil
	}

	return ButtonReleased, nil
}

//func update() {
// get press, when release get duration between press and release
// if duration > 200 between presses then we are tapping the button, couunt the taps
// if duration between press and release > 500 then we are holding
// reset long duration on release
// https://github.com/adafruit/Adafruit_CircuitPython_Debouncer/blob/main/adafruit_debouncer.py#L210
//}
