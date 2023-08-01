// Package is31fl3733 provides a driver for the Lumissil IS31FL3733 matrix LED
// driver.
//
// Driver supports following layouts:
//   - any custom LED matrix layout
//   - Adafruit 16x12 CharliePlex LED Matrix FeatherWing (CharlieWing)
//     https://www.adafruit.com/product/3163
//
// Datasheet:
//    https://www.lumissil.com/assets/pdf/core/IS31FL3733_DS.pdf
//
// This driver inspired by Adafruit Python driver:
//    https://github.com/y4m-y4m/CircuitPython_IS31FL3733
//
package is31fl3733

import (
	"fmt"
	//"time"

	"machine"
	//"math"
)

// Device implements TinyGo driver for Lumissil IS31FL3733 matrix LED driver


type Device struct {
	//Address uint8
	Bus     *machine.I2C

	// Currently selected Page register (one of the frame registers or the
	// function register)
	SelectedPage uint8
	//pwm_pixels []byte
	//abm_pixels []byte
	pixels [192]byte
}

func New(bus *machine.I2C, SPage uint8) Device {
	return Device{
		Bus:     bus,
		SelectedPage: SPage,
	}
}

// Configure chip for operating as a LED matrix display
func (d *Device) Init() (err error) {

	backlightPower := machine.RGB_POWER
	backlightPower.Configure(machine.PinConfig{Mode: machine.PinOutput})
	backlightPower.High()


	err = d.SelectPage(FUNCTION_REGISTER)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}
	// reset
	err = d.WriteCommonReg(_RESET_REGISTER, 0x01)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}
	// set software shutdown mode
	err = d.WriteCommonReg(_CONFIGURATION_REGISTER, CONFIGURATION_NORMAL_OPERATION)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}
	// set pullup_resistor = RESISTOR_32K & pulldown_resistor = RESISTOR_32K
	err = d.WriteCommonReg(_PULLUP_RESISTOR_SELECTION_REGISTER, RESISTOR_32K)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}
	err = d.WriteCommonReg(_PULLDOWN_RESISTOR_SELECTION_REGISTER, RESISTOR_32K)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}

	// init pixels
	for i:=0;i<192;i++ {
		d.pixels[i] = 0x01
	}

	// set brightness
	err = d.WriteCommonReg(_CURRENT_CONTROL_REGISTER, 128)
	if err != nil {
		return fmt.Errorf("failed to init: %w", err)
	}

	//err = d.SetLEDState()
	//if err != nil {
	//	return fmt.Errorf("failed to init: %w", err)
	//}
	return nil
}

// selectPage selects Page register, can be:
// - PWM registers 
// - function register
func (d *Device) SelectPage(Page uint8) (err error) {

	err = d.Bus.WriteRegister(ADDRESS, PAGE_REGISTER_WRITE_LOCK, []byte{WRITE_LOCK_DISABLE_ONCE})
	if err != nil {
		return fmt.Errorf("failed to unlock page write: %w", err)
	}
	err = d.Bus.WriteRegister(ADDRESS, PAGE_REGISTER, []byte{Page})
	if err != nil {
		return fmt.Errorf("failed to select page %w: %w", Page, err)
	}
	if Page != d.SelectedPage {
		d.SelectedPage = Page
	}
	return nil
}

func (d *Device) WriteCommonReg(reg_addr uint8, data uint8) (err error) {
	err = d.Bus.WriteRegister(ADDRESS, reg_addr, []byte{data})
	if err != nil {
		return fmt.Errorf("failed to write register: %w", err)
	}
	return nil
}


// WritePagedReg selects the function register and writes data into it
func (d *Device) WritePagedReg(reg_addr uint8, data uint8) (err error) {
	err = d.SelectPage(reg_addr)
	if err != nil {
		return err
	}
	err = d.Bus.WriteRegister(ADDRESS, reg_addr, []byte{data})
	if err != nil {
		return fmt.Errorf("failed to write page register: %w", err)
	}
	return nil
}

func (d *Device) EnableAllPixels() error {
	d.SelectPage(LED_CONTROL_REGISTER)
	err := d.Bus.WriteRegister(ADDRESS, LED_ONOFF_REGISTER_START, []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255})
	if err != nil {
		return fmt.Errorf("failed to write all pixels: %w", err)
	}
	d.SelectPage(PWM_REGISTER)

	data := make([]byte, 16)
	for i := range data {
		data[i] = 255
	}

	for i:=0;i<12;i++ {
		err := d.Bus.WriteRegister(ADDRESS, uint8(i*16), data)
		if err != nil {
			return fmt.Errorf("failed to write all pixels: %w", err)
		}
	}
	for i:=0;i<192;i++ {
		d.pixels[i] = 255
	}
	return nil
}


func (d *Device) Set_PWM_Pixel(index int, value [3]uint8) (err error) {
	buffer := make([]byte, 192)
	row := index / 16
	col := index & 15
	offset := row * 48 + col
	buffer[offset] = value[1]
	buffer[offset + 16] = value[0]
	buffer[offset + 32] = value[2]
	if index ==56 {
		// No.61 and No.62 are under the space key.
		// 61
		buffer[157] = value[1]
		buffer[157 + 16] = value[0]
		buffer[157 + 32] = value[2]
		// 62
		buffer[158] = value[1]
		buffer[158 + 16] = value[0]
		buffer[158 + 32] = value[2]
	}
	err = d.SelectPage(PWM_REGISTER)
	if err != nil {
		return fmt.Errorf("failed to write page register: %w", err)
	}
	return d.Bus.WriteRegister(ADDRESS, 0x00, buffer)
}