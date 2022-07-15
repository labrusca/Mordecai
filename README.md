# Mordecai
Project Mordecai is a keyboard firmware for Makerdiary M60 Keyboard based on tinygo

## Notice
Put file folder "tinygo" in your tinygo path before building.

## build
Run:
```
tinygo build -o KB-m60.uf2 -tagets m60-keyboard main.go
```
Let the keyboard into Bootloader mode, and copy the file 'KB-m60.uf2' to keyboard's drive.

Enjoy!

### Test on tinygo @0.24.0