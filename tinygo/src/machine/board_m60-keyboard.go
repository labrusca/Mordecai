// +build m60_keyboard

/*

Copyright [2021] [labrusca]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package machine

const HasLowFrequencyCrystal = true

// LEDs on the m600-keyboard (nRF52840 m.2 dev board)
const (
	LED       Pin = LED_GREEN
	LED_GREEN Pin = P0_29
	LED_RED   Pin = P0_30
	LED_BLUE  Pin = P0_31
)

// QSPI pins (unused)
const (
	QSPI_SCK   = P1_11
	QSPI_CS    = P1_13
	QSPI_DATA0 = P1_10
	QSPI_DATA1 = P1_14
	QSPI_DATA2 = P1_15
	QSPI_DATA3 = P1_12
)

// LED I2C pins on the m600-keyboard (unused)
const (
	SDA_PIN = P1_05
	SCL_PIN = P1_06
)

const (
	POWER = P1_04
	INTERRUPT = P1_07
)

// board
const (
	//Row
	R1 = P0_05
	R2 = P0_06
	R3 = P0_07
	R4 = P0_08
	R5 = P1_09
	R6 = P1_08
	R7 = P0_12
	R8 = P0_11
	//Col
	C1 = P0_19
	C2 = P0_20
	C3 = P0_21
	C4 = P0_22
	C5 = P0_23
	C6 = P0_24
	C7 = P0_25
	C8 = P0_26
	BUTTON = P0_27
)

// Battery
const (
	CHARGING = P0_03
	VOLTAGE = P0_02
)

// USB CDC identifiers
const (
	usb_STRING_PRODUCT      = "Makerdiary M60 Keyboard"
	usb_STRING_MANUFACTURER = "Nordic Semiconductor ASA"
)


// UART pins
const (
	UART_TX_PIN Pin = P0_16
	UART_RX_PIN Pin = P0_15
)

// SPI pins (unused)
const (
	SPI0_SCK_PIN = P1_15
	SPI0_SDO_PIN = P1_13
	SPI0_SDI_PIN = P1_14
)

var (
	usb_VID uint16 = 0x2886
	usb_PID uint16 = 0xF002
)
