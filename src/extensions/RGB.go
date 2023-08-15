package extensions

import (
	"m60/is31fl3733"
	"m60/hardware"
)

type RGBLED struct {
	Obj    is31fl3733.Device
}

func (l RGBLED) WhilePress(k hardware.Key) error {
	err := l.Obj.Set_PWM_Pixel(k.Num, [3]uint8{255, 255, 255})
	if err != nil {
		return err
	}
	return nil
}

func (l RGBLED) WhileRelease(k hardware.Key) error {
	err := l.Obj.Set_PWM_Pixel(k.Num, [3]uint8{0, 0, 0})
	if err != nil {
		return err
	}
	return nil
}