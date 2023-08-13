package main

import (
	"fmt"
	"log"
	//"container/list"
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

	KeyNO keyboard.Keycode = 0x00
	KeyFN keyboard.Keycode = 0x00

	KEYMAP = [...][61]keyboard.Keycode{
		{
			keyboard.KeyEsc, keyboard.Key1, keyboard.Key2, keyboard.Key3, keyboard.Key4, keyboard.Key5, keyboard.Key6, keyboard.Key7, keyboard.Key8, keyboard.Key9, keyboard.Key0, keyboard.KeyMinus, keyboard.KeyEqual, keyboard.KeyBackspace,
			keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT, keyboard.KeyY, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash,
			keyboard.KeyCapsLock, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG, keyboard.KeyH, keyboard.KeyJ, keyboard.KeyK, keyboard.KeyL, keyboard.KeySemicolon, keyboard.KeyQuote, keyboard.KeyEnter,
			keyboard.KeyModifierLeftShift, keyboard.KeyZ, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB, keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyModifierRightShift,
			keyboard.KeyModifierLeftCtrl, keyboard.KeyModifierLeftAlt, keyboard.KeyModifierLeftGUI, keyboard.KeySpace, keyboard.KeyModifierRightAlt, keyboard.KeyMenu, KeyFN, keyboard.KeyModifierRightCtrl,
		},

		{
			keyboard.KeyTilde, keyboard.KeyF1, keyboard.KeyF2, keyboard.KeyF3, keyboard.KeyF4, keyboard.KeyF5, keyboard.KeyF6, keyboard.KeyF7, keyboard.KeyF8, keyboard.KeyF9, keyboard.KeyF10, keyboard.KeyF11, keyboard.KeyF12, keyboard.KeyDelete,
			KeyNO, KeyNO, keyboard.KeyUp, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyMediaVolumeDec, keyboard.KeyMediaVolumeInc, keyboard.KeyMediaMute,
			KeyNO, keyboard.KeyLeft, keyboard.KeyDown, keyboard.KeyRight, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, 0xE470, 0xE46F, KeyNO,
			KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyMediaPrevTrack, keyboard.KeyMediaNextTrack, keyboard.KeyMediaPlayPause, KeyNO,
			KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO,
		},
	}

	pressed_keys []int
)

func Scan() ([]int, []int, []int, error) {
	var new_keys []int
	var inter_pressed_keys []int
	released_keys := pressed_keys
	for c, _ := range COLS {
		COLS[c].High()
		for r, _ := range ROWS {
			if ROWS[r].Get() {
				key := r*len(COLS) + c
				inter_pressed_keys = append(inter_pressed_keys, key)
				deleteIndex := IsContain(released_keys, key)
				if deleteIndex != -1 {
					released_keys = append(released_keys[:deleteIndex], released_keys[(deleteIndex+1):]...)
				} else {
					new_keys = append(new_keys, key)
				}

			}
		}
		COLS[c].Low()
		time.Sleep(time.Millisecond * 2)
	}
	pressed_keys = inter_pressed_keys

	return pressed_keys, released_keys, new_keys, nil
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

func ChkErr(n int, e error) {
	if e != nil {
		led_err := machine.LED_RED
		led_err.Configure(machine.PinConfig{Mode: machine.PinOutput})
		led_err.Low()
		log.Print("line: ", n)
		log.Println(e)
		time.Sleep(time.Millisecond * 5000)
		//os.Exit(0)
		machine.EnterUF2Bootloader()
	}
}

func power_init() {
	power := machine.POWER_PULLUP_PIN
	power.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
}

func main() {
	//power_init()
	btn := machine.BUTTON
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)

	//for btn.Get() == true {
	//	continue
	//}

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


	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 255})
	time.Sleep(time.Millisecond * 500)
	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 255})
	time.Sleep(time.Millisecond * 500)
	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	for i, _ := range ROWS {
		ROWS[i].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
	for j, _ := range COLS {
		COLS[j].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	var kb = keyboard.New()
	var layer int
	layer = 0

	for {
		// Press the button for entering Bootloader
		if btn.Get() == false {
			log.Println("Enter Bootloader...")
			machine.EnterUF2Bootloader()
		}

		_, released_keys, new_keys, err := Scan()
		if err != nil {
			log.Println(err)
		}

		for _, key := range released_keys {
			keycode := KEYMAP[layer][IsContain(COORDS, key)]
			if keycode == KeyFN {
				rgb.Set_PWM_Pixel(key, [3]uint8{0, 0, 0})
				layer = 0
				log.Print("Layer 0\n")
				continue
			}
			err := rgb.Set_PWM_Pixel(key, [3]uint8{0, 0, 0})
			ChkErr(183, err)
			err = kb.Up(keycode)
			ChkErr(185, err)
			fmt.Print("Release: ")
			fmt.Print(key, keycode)
		}

		for _, key := range new_keys {
			keycode := KEYMAP[layer][IsContain(COORDS, key)]
			if keycode == KeyFN {
				rgb.Set_PWM_Pixel(key, [3]uint8{255, 255, 0})
				layer = 1
				log.Print("Layer 1\n")
				continue
			}
			err := rgb.Set_PWM_Pixel(key, [3]uint8{255, 255, 255})
			ChkErr(199, err)
			err = kb.Down(keycode)
			ChkErr(201, err)
			fmt.Print("Press: ")
			fmt.Print(key, keycode)
		}

	}
	kb.Release()
	log.Print(fmt.Errorf("Unknow Error!"))
	machine.EnterUF2Bootloader()
}
