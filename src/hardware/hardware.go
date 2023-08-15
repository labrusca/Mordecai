package hardware

import (
	"m60/is31fl3733"
	"machine"
	"machine/usb/hid/keyboard"
	//"os"
	"time"
)
var (
	ROWS = [8]machine.Pin{machine.P0_05, machine.P0_06, machine.P0_07, machine.P0_08, machine.P1_09, machine.P1_08, machine.P0_12, machine.P0_11}
	COLS = [8]machine.Pin{machine.P0_19, machine.P0_20, machine.P0_21, machine.P0_22, machine.P0_23, machine.P0_24, machine.P0_25, machine.P0_26}

	COORDS = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14,
		28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41,
		53, 54, 55, 56, 57, 58, 59, 60,
	}

)

func InitPower(p machine.Pin) {
	p.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
}

func OpenUART() {
	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)
}

func InitRGB() is31fl3733.Device {
	rgb := is31fl3733.New(machine.I2C0, is31fl3733.FUNCTION_REGISTER)
	err := rgb.Bus.Configure(machine.I2CConfig{
		Mode:      machine.I2CModeController,
		Frequency: 400000,
		SCL:       machine.SCL_PIN,
		SDA:       machine.SDA_PIN,
	})
	if err != nil {
		println("could not configure I2C:", err)
	}

	err = rgb.Init()
	println("RGB init done!")
	if err != nil {
		println("could not init:", err)
	}
	err = rgb.EnableAllPixels()
	if err != nil {
		println("could not open lights:", err)
	}
	println("RGB all open!")
	return rgb
}

// https://blog.csdn.net/m0_37422289/article/details/103570799 CC 4.0 BY-SA
func IsContain(items []int, item int) int {
	for i, eachItem := range items {
		if eachItem == item {
			return i
		}
	}
	return -1
}



type Device struct {
	History [25]keyboard.Keycode
	Keys    [61]Key
	cols     [8]machine.Pin
	rows     [8]machine.Pin
	Layer   int
}

type Key struct{
	Num         int
	Ispress     bool
	NeedChange  bool
	Keycode     keyboard.Keycode
}

func NewDevice(r [8]machine.Pin, c [8]machine.Pin) Device {
	var k [61]Key
	for n ,c := range COORDS{
		k[n].Num = c
		k[n].Ispress = false
		k[n].NeedChange = false
	}
	for i, _ := range r {
		r[i].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
	for j, _ := range c {
		c[j].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	return Device {
		cols:  c,
		rows:  r,
		Keys: k,
	}
}

func (d *Device)UpdateStatus() error {

	for i, _ := range d.cols {
		d.cols[i].High()
		for j, _ := range d.rows {
			pressed := d.rows[j].Get()
			key := j * len(d.cols) + i
			keyIndex := IsContain(COORDS, key)
			if keyIndex == -1{
				continue
			}
			if d.Keys[keyIndex].Ispress != pressed {
				d.Keys[keyIndex].NeedChange  = true
			} else {
				d.Keys[keyIndex].NeedChange  = false
			}
			d.Keys[keyIndex].Ispress = pressed
		}
		d.cols[i].Low()
		time.Sleep(time.Millisecond * 2)
	}
	return nil
}