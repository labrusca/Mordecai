# Mordecai
Project Mordecai is a keyboard firmware for [Makerdiary M60 Keyboard](https://wiki.makerdiary.com/m60/) based on [TinyGo](https://github.com/tinygo-org/tinygo)

## Plan
- [x] Basic function of the USB Keyboard
- [x] Custom Extensions
- [ ] Bluetooth Support
- [x] Add is31fl3733 Driver as extension (Only PWM mode for now)
- [ ] Add MX25R6435F Driver
- [ ] USB Mass Storage Device Driver Support
- [ ] NFC Support
- [ ] Rewrite Golang code

## Notice
**First, you need add some files to your TINYGOROOT:**  
Put file folder "tinygo" in your tinygo path before building.

## Build
Run:
```
tinygo build -o KB-m60.uf2 -tagets m60-keyboard main.go
```
Let the keyboard into Bootloader mode, and copy the file 'KB-m60.uf2' to keyboard's drive.

Enjoy!

### Tested on TinyGo@0.28.1 & Golang@1.20.5