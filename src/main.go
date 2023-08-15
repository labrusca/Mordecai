package main

import (
	"fmt"
	"log"
	"m60/hardware"
	"m60/extensions"
	"machine"
	"machine/usb/hid/keyboard"
	//"os"
	"time"
)

var KEYMAP = [...][61]keyboard.Keycode{
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
		KeyNO, keyboard.KeyLeft, keyboard.KeyDown, keyboard.KeyRight, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyF21, keyboard.KeyF20, KeyNO,
		KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, keyboard.KeyMediaPrevTrack, keyboard.KeyMediaNextTrack, keyboard.KeyMediaPlayPause, KeyNO,
		KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO, KeyNO,
	},
}

var KeyNO keyboard.Keycode = 0x00
var KeyFN keyboard.Keycode = 0x00

func main() {
	//power_init()
	btn := machine.BUTTON
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	// For Debug
	//for btn.Get() == true {
	//	continue
	//}

	hardware.InitPower(machine.POWER_PULLUP_PIN)
	hardware.OpenUART()





	kb := keyboard.New()
	m60 := hardware.NewDevice(hardware.ROWS, hardware.COLS)
	rgb := hardware.InitRGB()

	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 255})
	time.Sleep(time.Millisecond * 500)
	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 255})
	time.Sleep(time.Millisecond * 500)
	rgb.Set_PWM_Pixel(63, [3]uint8{0, 0, 0})
	time.Sleep(time.Millisecond * 500)

	//var exten_rgb extensions.Extensions
	exten_rgb := new(extensions.RGBLED)
	exten_rgb.Obj = rgb


	for {
		// Press the button for entering Bootloader
		if btn.Get() == false {
			log.Println("Enter Bootloader...")
			machine.EnterUF2Bootloader()
		}

		err := m60.UpdateStatus()
		if err != nil {
			print(fmt.Errorf("Unknow Error!"))
			break
		}

		for i:=0; i<61; i++ {
			m60.Keys[i].Keycode = KEYMAP[m60.Layer][i]
			if !m60.Keys[i].NeedChange {
				continue
			}

			if m60.Keys[i].Ispress {
				//rgb.Set_PWM_Pixel(m60.Keys[i].Num, [3]uint8{255, 255, 255})
				exten_rgb.WhilePress(m60.Keys[i])
				if m60.Keys[i].Keycode == KeyFN {
					m60.Layer = 1
					log.Println("Layer 1\n")
					continue
				}
				kb.Down(m60.Keys[i].Keycode)
				println("Press: ",m60.Keys[i].Num, m60.Keys[i].Keycode)
				fmt.Print()
			} else {
				//rgb.Set_PWM_Pixel(m60.Keys[i].Num, [3]uint8{0, 0, 0})
				exten_rgb.WhileRelease(m60.Keys[i])
				if m60.Keys[i].Keycode == KeyFN {
					m60.Layer = 0
					log.Println("Layer 0\n")
					kb.Release()
					continue
				}
				kb.Up(m60.Keys[i].Keycode)
				println("Release: ",m60.Keys[i].Num, m60.Keys[i].Keycode)
			}
			m60.Keys[i].NeedChange = false
		}

	}


	kb.Release()
	log.Println(fmt.Errorf("Unknow Error!"))
	machine.EnterUF2Bootloader()
}
