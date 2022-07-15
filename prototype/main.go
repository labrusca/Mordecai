package main

import (
    "machine"
	"machine/usb/hid/keyboard"
    "time"
)


var (
	ROWS = [8]machine.Pin{machine.P0_05, machine.P0_06, machine.P0_07, machine.P0_08, machine.P1_09, machine.P1_08, machine.P0_12, machine.P0_11}
	COLS = [8]machine.Pin{machine.P0_19, machine.P0_20, machine.P0_21, machine.P0_22, machine.P0_23, machine.P0_24, machine.P0_25, machine.P0_26}

	COORDS = []int{
		0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13,
		27,26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14,
		28,29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,     40,
		52,51, 50, 49, 48, 47, 46, 45, 44, 43, 42,         41,
		53,  54, 55,             56,           57, 58, 59, 60,
	}

	KeyNO keyboard.Keycode = 0x00
	KeyFN keyboard.Keycode = 0xC0

	KEYMAP = [...][61]keyboard.Keycode{
		{
			keyboard.KeyEsc, keyboard.Key1,   keyboard.Key2,   keyboard.Key3,   keyboard.Key4,   keyboard.Key5,   keyboard.Key6,   keyboard.Key7,   keyboard.Key8,   keyboard.Key9,   keyboard.Key0,   keyboard.KeyMinus, keyboard.KeyEqual,  keyboard.KeyBackspace,
			keyboard.KeyTab,  keyboard.KeyQ,    keyboard.KeyW,    keyboard.KeyE,    keyboard.KeyR,    keyboard.KeyT,    keyboard.KeyY,    keyboard.KeyU,    keyboard.KeyI,    keyboard.KeyO,    keyboard.KeyP,    keyboard.KeyLeftBrace, keyboard.KeyRightBrace, keyboard.KeyBackslash,
			keyboard.KeyCapsLock, keyboard.KeyA,    keyboard.KeyS,    keyboard.KeyD,    keyboard.KeyF,    keyboard.KeyG,    keyboard.KeyH,    keyboard.KeyJ,    keyboard.KeyK,    keyboard.KeyL,    keyboard.KeySemicolon, keyboard.KeyQuote,           keyboard.KeyEnter,
			keyboard.KeyModifierLeftShift, keyboard.KeyZ,    keyboard.KeyX,    keyboard.KeyC,    keyboard.KeyV,    keyboard.KeyB,    keyboard.KeyN,    keyboard.KeyM,    keyboard.KeyComma, keyboard.KeyPeriod,  keyboard.KeySlash,   keyboard.KeyModifierRightShift,
			keyboard.KeyModifierLeftCtrl, keyboard.KeyModifierLeftAlt, keyboard.KeyModifierLeftGUI,                       keyboard.KeySpace,                     keyboard.KeyModifierRightAlt, keyboard.KeyModifierRightGUI, KeyFN, keyboard.KeyModifierRightCtrl,
		},

		{
			keyboard.KeyEsc,    keyboard.KeyF1,   keyboard.KeyF2,   keyboard.KeyF3,   keyboard.KeyF4,   keyboard.KeyF5,   keyboard.KeyF6,   keyboard.KeyF7,   keyboard.KeyF8,   keyboard.KeyF9,   keyboard.KeyF10,  keyboard.KeyF11,  keyboard.KeyF12,  keyboard.KeyDelete,
			KeyNO, KeyNO, keyboard.KeyUp,   KeyNO,   KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyMediaVolumeDec, keyboard.KeyMediaVolumeInc, keyboard.KeyMediaMute,
			KeyNO, keyboard.KeyLeft, keyboard.KeyDown, keyboard.KeyRight, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO,
			KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyMediaPrevTrack, keyboard.KeyMediaNextTrack, keyboard.KeyMediaPlayPause, KeyNO,
			KeyNO, KeyNO, KeyNO, KeyNO, KeyNO,          KeyNO, KeyNO, KeyNO,
		},
	}

	pressed_keys []int
)

func Scan() ([]int, []int, []int) {
	var new_keys []int
	var inter_pressed_keys []int
	released_keys := pressed_keys
	for c, _ :=  range COLS{
		COLS[c].High()
		for r, _ := range ROWS{
			if ROWS[r].Get() {
				key := r * len(COLS) + c
				inter_pressed_keys = append(inter_pressed_keys, key)
				deleteIndex := IsContain(released_keys,key)
				if deleteIndex != -1 {
					released_keys = append(released_keys[:deleteIndex], released_keys[(deleteIndex+1):]...)
				} else {
					new_keys = append(new_keys, key)
				}

			}
		}
		COLS[c].Low()
	}
	pressed_keys = inter_pressed_keys
	return pressed_keys, released_keys, new_keys
}

//https://blog.csdn.net/m0_37422289/article/details/103570799 CC 4.0 BY-SA
func IsContain(items []int, item int) int {
	for i, eachItem := range items {
		if eachItem == item {
			return  i
		}
	}
	return -1
}

func main() {
	btn := machine.BUTTON
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	for i, _ := range ROWS {
		ROWS[i].Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	}
	for j, _ := range COLS {
		COLS[j].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	config := machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	}
	machine.UART0.Configure(config)

	var keyboard = keyboard.New()
	var layer int
	layer = 0
	for {
		// Press the button for entering Bootloader
		if btn.Get() == false{
            machine.EnterUF2Bootloader()
        }
		_, released_keys, new_keys := Scan()

		for _, key := range released_keys {
			if KEYMAP[layer][IsContain(COORDS,key)] == KeyFN {
				layer = 0
				println("Layer 0")
				continue
			}
			print("Release:")
			println(key, KEYMAP[layer][IsContain(COORDS,key)])
			keyboard.Up(KEYMAP[layer][IsContain(COORDS,key)])
		}

		for _, key := range new_keys {
			if KEYMAP[layer][IsContain(COORDS,key)] == KeyFN {
				layer = 1
				print("Layer 1")
				continue
			}
			println("Press:")
			println(key, KEYMAP[layer][IsContain(COORDS,key)])
			keyboard.Down(KEYMAP[layer][IsContain(COORDS,key)])
		}
		time.Sleep(time.Millisecond * 1)
	}
}