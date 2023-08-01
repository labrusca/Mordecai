# Mordecai
Project Mordecai is a keyboard firmware for Makerdiary M60 Keyboard based on tinygo

## Plan
- [x] Base Keyboard function
- [x] Add is31fl3733 driver(Only PWM mode)
- [ ] Bluetooth Support
- [ ] Rewrite Golang code

## Notice
**First, you need add some files to your TINYGOROOT**
Put file folder "tinygo" in your tinygo path before building.

## build
Run:
```
tinygo build -o KB-m60.uf2 -tagets m60-keyboard main.go
```
Let the keyboard into Bootloader mode, and copy the file 'KB-m60.uf2' to keyboard's drive.

Enjoy!

### Test on tinygo @0.28.1