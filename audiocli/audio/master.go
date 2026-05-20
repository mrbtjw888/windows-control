package audio

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

func ChangeMasterVolume(val string) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Create device enumerator
	unk, err := ole.CreateInstance(
		wca.CLSID_MMDeviceEnumerator,
		wca.IID_IMMDeviceEnumerator,
	)
	if err != nil {
		fmt.Println("CreateInstance failed:", err)
		return
	}

	enumerator := (*wca.IMMDeviceEnumerator)(unsafe.Pointer(unk))

	// Get default audio device
	var device *wca.IMMDevice
	enumerator.GetDefaultAudioEndpoint(
		wca.ERender,
		wca.EConsole,
		&device,
	)

	// Activate endpoint volume interface
	var endpoint *wca.IAudioEndpointVolume
	device.Activate(
		wca.IID_IAudioEndpointVolume,
		ole.CLSCTX_ALL,
		nil,
		(*unsafe.Pointer)(unsafe.Pointer(&endpoint)),
	)

	// Get current volume
	var current float32
	endpoint.GetMasterVolumeLevelScalar(&current)

	newVol := computeMasterVolume(val, current)

	endpoint.SetMasterVolumeLevelScalar(newVol, nil)

	fmt.Printf("Master volume: %.2f → %.2f\n", current, newVol)
}

func SetMasterMute(mute bool) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unk, _ := ole.CreateInstance(
		wca.CLSID_MMDeviceEnumerator,
		wca.IID_IMMDeviceEnumerator,
	)

	enumerator := (*wca.IMMDeviceEnumerator)(unsafe.Pointer(unk))

	var device *wca.IMMDevice
	enumerator.GetDefaultAudioEndpoint(
		wca.ERender,
		wca.EConsole,
		&device,
	)

	var endpoint *wca.IAudioEndpointVolume
	device.Activate(
		wca.IID_IAudioEndpointVolume,
		ole.CLSCTX_ALL,
		nil,
		(*unsafe.Pointer)(unsafe.Pointer(&endpoint)),
	)

	endpoint.SetMute(mute, nil)

	fmt.Println("Master mute:", mute)
}

func computeMasterVolume(input string, current float32) float32 {
	if strings.HasPrefix(input, "+") || strings.HasPrefix(input, "-") {
		delta, _ := strconv.ParseFloat(input, 32)
		return clamp(current + float32(delta)/100.0)
	}

	val, _ := strconv.ParseFloat(input, 32)
	return clamp(float32(val) / 100.0)
}