package is31fl3733

// Registers. Names taken from the datasheet:
// https://www.lumissil.com/assets/pdf/core/IS31FL3731_DS.pdf
const (
	// AD pin connected to GND
	//I2C_ADDRESS_74 uint8 = 0x74
	// AD pin connected to SCL
	//I2C_ADDRESS_75 uint8 = 0x75
	// AD pin connected to SDA
	//I2C_ADDRESS_76 uint8 = 0x76
	// AD pin connected to VCC
	//I2C_ADDRESS_77 uint8 = 0x77

	ADDRESS uint8 = 0x50

	// REGISTER DEFINITION
	PAGE_REGISTER            uint8 = 0xFD
	PAGE_REGISTER_WRITE_LOCK uint8 = 0xFE
	INTERRUPT_MASK_REGISTER     uint8 = 0xF0
	INTERRUPT_STATUS_REGISTER   uint8 = 0xF1
	LED_CONTROL_REGISTER        uint8 = 0x00
	PWM_REGISTER                uint8 = 0x01
	AUTO_BREATH_MODE_REGISTER   uint8 = 0x02
	FUNCTION_REGISTER           uint8 = 0x03

	WRITE_LOCK_DISABLE_ONCE  uint8 = 0xC5
	LED_ONOFF_REGISTER_START uint8 = 0x00
	LED_OPEN_REGISTER_START  uint8 = 0x18
	LED_SHORT_REGISTER_START uint8 = 0x30

	_CONFIGURATION_REGISTER               uint8 = 0x00
	_CURRENT_CONTROL_REGISTER             uint8 = 0x01
	_TIME_UPDATE_REGISTER                 uint8 = 0x0E
	_PULLUP_RESISTOR_SELECTION_REGISTER   uint8 = 0x0F
	_PULLDOWN_RESISTOR_SELECTION_REGISTER uint8 = 0x10
	_RESET_REGISTER                       uint8 = 0x11

	CONFIGURATION_SYNC_CLOCK_MASTER            uint8 = 0b01000000
	CONFIGURATION_SYNC_CLOCK_SLAVE             uint8 = 0b10000000
	CONFIGURATION_SYNC_CLOCK_HIGH_IMPEDANCE    uint8 = 0b00000000
	CONFIGURATION_OPEN_SHORT_DETECTION_ENABLE  uint8 = 0b00000100
	CONFIGURATION_OPEN_SHORT_DETECTION_DISABLE uint8 = 0b00000000
	CONFIGURATION_ABM_ENABLE                   uint8 = 0b00000010
	CONFIGURATION_PWM_ENABLE                   uint8 = 0b00000000
	CONFIGURATION_SOFTWARE_SHUTDOWN            uint8 = 0b00000000
	CONFIGURATION_NORMAL_OPERATION             uint8 = 0b00000001

	AUTO_CLEAR_INTERRUPT_ENABLE uint8 = 0x08
	AUTO_CLEAR_INTERRUPT_DISABLE uint8 = 0x00
	AUTO_BREATH_INTERRUPT_ENABLE uint8 = 0x04
	AUTO_BREATH_INTERRUPT_DISABLE uint8 = 0x00
	DOT_SHORT_INTERRUPT_ENABLE uint8 = 0x02
	DOT_SHORT_INTERRUPT_DISABLE uint8 = 0x00
	DOT_OPEN_INTERRUPT_ENABLE uint8 = 0x01
	DOT_OPEN_INTERRUPT_DISABLE uint8 = 0x00

	LED_MODE_PWM  uint8 = 0x00
	LED_MODE_ABM1 uint8 = 0x01
	LED_MODE_ABM2 uint8 = 0x02
	LED_MODE_ABM3 uint8 = 0x03

	ABM_T1_T3_210MS   uint8 = 0x00
	ABM_T1_T3_420MS   uint8 = 0x20
	ABM_T1_T3_840MS   uint8 = 0x40
	ABM_T1_T3_1680MS  uint8 = 0x60
	ABM_T1_T3_3360MS  uint8 = 0x80
	ABM_T1_T3_6720MS  uint8 = 0xA0
	ABM_T1_T3_13440MS uint8 = 0xC0
	ABM_T1_T3_26880MS uint8 = 0xE0

	ABM_T2_T4_0MS     uint8 = 0x00
	ABM_T2_T4_210MS   uint8 = 0x02
	ABM_T2_T4_420MS   uint8 = 0x04
	ABM_T2_T4_840MS   uint8 = 0x06
	ABM_T2_T4_1680MS  uint8 = 0x08
	ABM_T2_T4_3360MS  uint8 = 0x0A
	ABM_T2_T4_6720MS  uint8 = 0x0C
	ABM_T2_T4_13440MS uint8 = 0x0E
	ABM_T2_T4_26880MS uint8 = 0x10

	ABM_T4_53760MS  uint8 = 0x12
	ABM_T4_107520MS uint8 = 0x14

	ABM_LOOP_BEGIN_T1 uint8 = 0x00
	ABM_LOOP_BEGIN_T2 uint8 = 0x10
	ABM_LOOP_BEGIN_T3 uint8 = 0x20
	ABM_LOOP_BEGIN_T4 uint8 = 0x30

	ABM_LOOP_END_T3 uint8 = 0x00
	ABM_LOOP_END_T1 uint8 = 0x40

	RESISTOR_NONE uint8 = 0x00
	RESISTOR_500  uint8 = 0x01
	RESISTOR_1K   uint8 = 0x02
	RESISTOR_2K   uint8 = 0x03
	RESISTOR_4K   uint8 = 0x04
	RESISTOR_8K   uint8 = 0x05
	RESISTOR_16K  uint8 = 0x06
	RESISTOR_32K  uint8 = 0x07
)