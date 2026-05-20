package main

import (
	"fmt"
	"os"
	"strings"

	"audiocli/audio"
	"audiocli/media"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := strings.ToLower(os.Args[1])

	switch cmd {

	case "list":
		audio.ListSessions()

	case "volume":
		if len(os.Args) < 4 {
			fmt.Println("Usage: volume <process> <value|+delta|-delta>")
			return
		}
		audio.ChangeVolume(os.Args[2], os.Args[3])

	case "mute":
		if len(os.Args) < 3 {
			fmt.Println("Usage: mute <process>")
			return
		}
		audio.SetMute(os.Args[2], true)

	case "unmute":
		if len(os.Args) < 3 {
			fmt.Println("Usage: unmute <process>")
			return
		}
		audio.SetMute(os.Args[2], false)

	case "control":
		if len(os.Args) < 3 {
			fmt.Println("Usage: control <playpause|next|prev|volup|voldown|mute>")
			return
		}
		media.Handle(strings.ToLower(os.Args[2]))

	case "master":
		if len(os.Args) < 3 {
			fmt.Println("Usage: master <value|+delta|-delta>")
			return
		}
		audio.ChangeMasterVolume(os.Args[2])

	case "mastermute":
		audio.SetMasterMute(true)

	case "masterunmute":
		audio.SetMasterMute(false)

	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("audiocli - control system audio")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  list")
	fmt.Println("      List active audio sessions")
	fmt.Println()
	fmt.Println("  volume <process> <value>")
	fmt.Println("      Set volume (0-100)")
	fmt.Println("      Or relative: +10 / -10")
	fmt.Println()
	fmt.Println("  mute <process>")
	fmt.Println("  unmute <process>")
	fmt.Println()
	fmt.Println("  control <action>")
	fmt.Println("      playpause | next | prev | volup | voldown | mute")
	fmt.Println("  master <value>")
	fmt.Println("      Set volume (0-100)")
	fmt.Println("      Or relative: +10 / -10")
}

// go build -ldflags="-s -w -H=windowsgui" -o audio.exe