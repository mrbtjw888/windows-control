package media

import (
	"golang.org/x/sys/windows"
)

var (
	user32 = windows.NewLazySystemDLL("user32.dll")
	keybd  = user32.NewProc("keybd_event")
)

func press(code byte) {
	keybd.Call(uintptr(code), 0, 0, 0)
}

func Handle(action string) {
	switch action {

	case "playpause":
		press(0xB3)

	case "next":
		press(0xB0)

	case "prev":
		press(0xB1)

	case "mute":
		press(0xAD)

	case "volup":
		press(0xAF)

	case "voldown":
		press(0xAE)
	}
}