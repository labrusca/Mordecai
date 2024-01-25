package main

import (
	"context"
	_ "embed"
	"log"
	"m60/is31fl3733"
	"machine"
	"machine/usb"
	kc "machine/usb/hid/keyboard"
	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
	"time"
)


var (
	ROWS = []machine.Pin{machine.P0_05, machine.P0_06, machine.P0_07, machine.P0_08, machine.P1_09, machine.P1_08, machine.P0_12, machine.P0_11}
	COLS = []machine.Pin{machine.P0_19, machine.P0_20, machine.P0_21, machine.P0_22, machine.P0_23, machine.P0_24, machine.P0_25, machine.P0_26}
	KeyBrightnessUp   = keyboard.Keycode(0xE400 | 0x6F)
	KeyBrightnessDown = keyboard.Keycode(0xE400 | 0x70)
	KEYMAP = [][]keyboard.Keycode{
		{
			keyboard.Keycode(kc.KeyEsc), keyboard.Keycode(kc.Key1), keyboard.Keycode(kc.Key2), keyboard.Keycode(kc.Key3), keyboard.Keycode(kc.Key4), keyboard.Keycode(kc.Key5), keyboard.Keycode(kc.Key6), keyboard.Keycode(kc.Key7), keyboard.Keycode(kc.Key8), keyboard.Keycode(kc.Key9), keyboard.Keycode(kc.Key0), keyboard.Keycode(kc.KeyMinus), keyboard.Keycode(kc.KeyEqual), keyboard.Keycode(kc.KeyBackspace),
			keyboard.Keycode(kc.KeyTab), keyboard.Keycode(kc.KeyQ), keyboard.Keycode(kc.KeyW), keyboard.Keycode(kc.KeyE), keyboard.Keycode(kc.KeyR), keyboard.Keycode(kc.KeyT), keyboard.Keycode(kc.KeyY), keyboard.Keycode(kc.KeyU), keyboard.Keycode(kc.KeyI), keyboard.Keycode(kc.KeyO), keyboard.Keycode(kc.KeyP), keyboard.Keycode(kc.KeyLeftBrace), keyboard.Keycode(kc.KeyRightBrace), keyboard.Keycode(kc.KeyBackslash),
			keyboard.Keycode(kc.KeyCapsLock), keyboard.Keycode(kc.KeyA), keyboard.Keycode(kc.KeyS), keyboard.Keycode(kc.KeyD), keyboard.Keycode(kc.KeyF), keyboard.Keycode(kc.KeyG), keyboard.Keycode(kc.KeyH), keyboard.Keycode(kc.KeyJ), keyboard.Keycode(kc.KeyK), keyboard.Keycode(kc.KeyL), keyboard.Keycode(kc.KeySemicolon), keyboard.Keycode(kc.KeyQuote), keyboard.Keycode(kc.KeyEnter),
			keyboard.Keycode(kc.KeyModifierLeftShift), keyboard.Keycode(kc.KeyZ), keyboard.Keycode(kc.KeyX), keyboard.Keycode(kc.KeyC), keyboard.Keycode(kc.KeyV), keyboard.Keycode(kc.KeyB), keyboard.Keycode(kc.KeyN), keyboard.Keycode(kc.KeyM), keyboard.Keycode(kc.KeyComma), keyboard.Keycode(kc.KeyPeriod), keyboard.Keycode(kc.KeySlash), keyboard.Keycode(kc.KeyModifierRightShift),
			keyboard.Keycode(kc.KeyModifierLeftCtrl), keyboard.Keycode(kc.KeyModifierLeftAlt), keyboard.Keycode(kc.KeyModifierLeftGUI), keyboard.Keycode(kc.KeySpace), keyboard.Keycode(kc.KeyModifierRightAlt), keyboard.Keycode(kc.KeyMenu), jp.KeyMod1, keyboard.Keycode(kc.KeyModifierRightCtrl),
		},

		{
			jp.KeyHankaku, jp.KeyF1, jp.KeyF2, jp.KeyF3, jp.KeyF4, jp.KeyF5, jp.KeyF6, jp.KeyF7, jp.KeyF8, jp.KeyF9, jp.KeyF10, jp.KeyF11, jp.KeyF12, jp.KeyDelete,
			jp.KeyTo0, jp.KeyHome, jp.KeyTo0, jp.KeyEnd, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyMediaVolumeDec, jp.KeyMediaVolumeInc, jp.KeyMediaMute,
			jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyLeft, jp.KeyDown, jp.KeyUp, jp.KeyRight, KeyBrightnessDown, KeyBrightnessUp, jp.KeyTo0,
			jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyMediaPrevTrack, jp.KeyMediaNextTrack, jp.KeyMediaPlayPause, jp.KeyTo0,
			jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0, jp.KeyTo0,
		},
	}

	COORDS = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14,
		28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41,
		53, 54, 55, 56, 57, 58, 59, 60,
	}
)

//Init RGB LED
var led = initRGB()

func main() {
	//machine.Flash.EraseBlocks(0, 1)
	usb.Product = "M60-Keyboard-0.1.0"

	//power_init()
	btn := machine.BUTTON
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	//openUART()

	led.Set_PWM_Pixel(63, [3]uint8{255, 0, 0})
	time.Sleep(time.Millisecond * 500)
	led.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	led.Set_PWM_Pixel(63, [3]uint8{255, 0, 0})
	time.Sleep(time.Millisecond * 500)
	led.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	
	err := run()
	if err != nil {
		log.Fatal(err)
		log.Println("Enter Bootloader...")
		machine.EnterUF2Bootloader()
	}
}

func run() error {
	d := keyboard.New()

	kboard := d.AddMatrixKeyboard(COLS, ROWS, KEYMAP)

	kboard.SetCallback(func(layer, index int, state keyboard.State) {
		var colour[3]uint8
		if layer == 0 {
			colour = [3]uint8{255, 255, 255}
		} else {
			colour = [3]uint8{255, 0, 0}
		}
		if state == keyboard.Press {
			led.Set_PWM_Pixel(index, colour)
		} else {
			led.Set_PWM_Pixel(index, [3]uint8{0, 0, 0})
		}
		log.Printf("rk: %d %d %d\n", layer, index, state)
		//changed.Set(1)
	})
	//fix for M60 keyboard
	for i:=0; i<61; i++ {
		for f := range KEYMAP{
			kboard.SetKeycode(f, COORDS[i], KEYMAP[f][i])
		}
	}
	// for Vial
	loadKeyboardDef()
	keyboard.Save()
	
	err := d.Init()
	if err != nil {
		return err
	}

	d.Debug = true
	return d.Loop(context.Background())
}

func openUART() {
	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)
}

func initRGB() is31fl3733.Device {
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
