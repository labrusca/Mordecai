package main

import (
    "machine"
	"machine/usb"
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

	KEYMAP = [...][61]usb.Keycode{
		{
			usb.KeyEsc, usb.Key1,   usb.Key2,   usb.Key3,   usb.Key4,   usb.Key5,   usb.Key6,   usb.Key7,   usb.Key8,   usb.Key9,   usb.Ke0,   usb.KeyMinus, usb.KeyEqual,  usb.KeyBackspace,
			usb.KeyTab,  usb.KeyQ,    usb.KeyW,    usb.KeyE,    usb.KeyR,    usb.KeyT,    usb.KeyY,    usb.KeyU,    usb.KeyI,    usb.KeyO,    usb.KeyP,    usb.KeyLeftBrace, usb.KeyRightBrace, usb.KeyBackslash,
			usb.KeyCapsLock, usb.KeyA,    usb.KeyS,    usb.KeyD,    usb.KeyF,    usb.KeyG,    usb.KeyH,    usb.KeyJ,    usb.KeyK,    usb.KeyL,    usb.KeySemicolon, usb.KeyQuote,           usb.KeyEnter,
			usb.KeyModifierLeftShift, usb.KeyZ,    usb.KeyX,    usb.KeyC,    usb.KeyV,    usb.KeyB,    usb.KeyN,    usb.KeyM,    usb.KeyComma, usb.KeyPeriod,  usb.KeySlash,   usb.KeyModifierRightShift,
			usb.KeyModifierLeftCtrl, usb.KeyModifierLeftAlt, usb.KeyModifierLeftGUI,                       usb.KeySpace,                     usb.KeyModifierRightAlt, usb.KeyModifierRightGUI, usb.KeyFN1, usb.KeyModifierRightCtrl,
		},

		{
			usb.KeyEsc,    usb.KeyF1,   usb.KeyF2,   usb.KeyF3,   usb.KeyF4,   usb.KeyF5,   usb.KeyF6,   usb.KeyF7,   usb.KeyF8,   usb.KeyF9,   usb.KeyF10,  usb.KeyF11,  usb.KeyF12,  usb.KeyDelete,
			usb.KeyNO, usb.KeyNO, usb.KeyUp,   usb.KeyNO,   usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyMediaVolumeDec, usb.KeyMediaVolumeInc, usb.KeyMediaMute,
			usb.KeyNO, usb.KeyLeft, usb.KeyDown, usb.KeyRight, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO,
			usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyBOOTLOADER, usb.KeyNO, usb.KeyNO, usb.KeyMediaPrevTrack, usb.KeyMediaNextTrack, usb.KeyMediaPlayPause, usb.KeyNO,
			usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO, usb.KeyNO,          usb.KeyNO, usb.KeyNO, usb.KeyNO,
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
	var keyboard = machine.HID0.Keyboard()



	for {
		// Press the button for entering Bootloader
		if btn.Get() == false{
            machine.EnterUF2Bootloader()
        }
		_, released_keys, new_keys := Scan()

		for _, key := range released_keys {
			print("Release:")
			println(key,KEYMAP[0][IsContain(COORDS,key)])
			keyboard.Up(KEYMAP[0][IsContain(COORDS,key)])
		}

		for _, key := range new_keys {
			print("Press:")
			println(key,KEYMAP[0][IsContain(COORDS,key)])
			keyboard.Down(KEYMAP[0][IsContain(COORDS,key)])
		}
		time.Sleep(time.Millisecond * 1)
	}
}