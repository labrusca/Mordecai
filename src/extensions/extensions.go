package extensions

import (
	"m60/hardware"
)

type Extensions interface {
	WhilePress(Key hardware.Key) error
	WhileRelease(Key hardware.Key) error
}