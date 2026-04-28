package main

import (
	"flag"
	"fmt"
	"syscall"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procKeybdEvent   = user32.NewProc("keybd_event")
)

// Windows Virtual Key Codes
const (
	VK_MEDIA_NEXT_TRACK = 0xB1
	VK_MEDIA_PREV_TRACK = 0xB0
	VK_MEDIA_PLAY_PAUSE = 0xB3
	VK_VOLUME_MUTE      = 0xAD
	VK_VOLUME_DOWN      = 0xAE
	VK_VOLUME_UP        = 0xAF
)

// sendKey simulates a physical key press
func sendKey(vk byte) {
	procKeybdEvent.Call(uintptr(vk), 0, 0, 0)         // Key Down
	procKeybdEvent.Call(uintptr(vk), 0, 0x0002, 0)    // Key Up
}

func main() {
	// 1. Setup Flags
	playPause := flag.Bool("playPause", false, "Toggle Play/Pause")
	next      := flag.Bool("next", false, "Next Track")
	prev      := flag.Bool("prev", false, "Previous Track")
	volUp     := flag.Bool("volUp", false, "Increase Volume")
	volDown   := flag.Bool("volDown", false, "Decrease Volume")
	mute      := flag.Bool("mute", false, "Toggle Mute")
	status    := flag.Bool("status", false, "Show current status")

	flag.Parse()

	// 2. Execute Actions
	actionsTaken := 0

	if *playPause { 
		sendKey(VK_MEDIA_PLAY_PAUSE)
		fmt.Println("Sent: Play/Pause")
		actionsTaken++
	}
	if *next { 
		sendKey(VK_MEDIA_NEXT_TRACK)
		fmt.Println("Sent: Next Track")
		actionsTaken++
	}
	if *prev { 
		sendKey(VK_MEDIA_PREV_TRACK)
		fmt.Println("Sent: Previous Track")
		actionsTaken++
	}
	if *volUp { 
		sendKey(VK_VOLUME_UP)
		fmt.Println("Sent: Volume Up")
		actionsTaken++
	}
	if *volDown { 
		sendKey(VK_VOLUME_DOWN)
		fmt.Println("Sent: Volume Down")
		actionsTaken++
	}
	if *mute { 
		sendKey(VK_VOLUME_MUTE)
		fmt.Println("Sent: Mute Toggle")
		actionsTaken++
	}

	// 3. Status logic
	if *status {
		fmt.Println("--- Windows Audio Controller ---")
		fmt.Println("Status: Listening for Global Media Sessions...")
		fmt.Println("Note: Use -help to see all available commands.")
		actionsTaken++
	}

	// If no flags were passed, show help
	if actionsTaken == 0 {
		flag.Usage()
	}
}

// compile: go build -ldflags="-s -w -H=windowsgui" -o audio.exe