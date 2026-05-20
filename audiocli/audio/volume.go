package audio

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

func ChangeVolume(target string, val string) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unk, _ := ole.CreateInstance(
		wca.CLSID_MMDeviceEnumerator,
		wca.IID_IMMDeviceEnumerator,
	)

	enumerator := (*wca.IMMDeviceEnumerator)(unsafe.Pointer(unk))

	var device *wca.IMMDevice
	enumerator.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &device)

	var manager *wca.IAudioSessionManager2
	device.Activate(
		wca.IID_IAudioSessionManager2,
		ole.CLSCTX_ALL,
		nil,
		(*unsafe.Pointer)(unsafe.Pointer(&manager)),
	)

	var sessionEnum *wca.IAudioSessionEnumerator
	manager.GetSessionEnumerator(&sessionEnum)

	var count int
	sessionEnum.GetCount(&count)

	for i := 0; i < count; i++ {

		var ctl *wca.IAudioSessionControl
		sessionEnum.GetSession(i, &ctl)

		unk2, _ := ctl.QueryInterface(wca.IID_IAudioSessionControl2)
		ctl2 := (*wca.IAudioSessionControl2)(unsafe.Pointer(unk2))

		var pid uint32
		ctl2.GetProcessId(&pid)

		name := getProcessName(pid)
		if !strings.Contains(strings.ToLower(name), strings.ToLower(target)) {
			continue
		}

		unkVol, _ := ctl.QueryInterface(wca.IID_ISimpleAudioVolume)
		vol := (*wca.ISimpleAudioVolume)(unsafe.Pointer(unkVol))

		var current float32
		vol.GetMasterVolume(&current)

		newVol := computeVolume(val, current)
		vol.SetMasterVolume(newVol, nil)

		fmt.Printf("%s: %.2f → %.2f\n", name, current, newVol)
	}
}

func computeVolume(input string, current float32) float32 {
	if strings.HasPrefix(input, "+") || strings.HasPrefix(input, "-") {
		delta, _ := strconv.ParseFloat(input, 32)
		return clamp(current + float32(delta)/100.0)
	}

	val, _ := strconv.ParseFloat(input, 32)
	return clamp(float32(val) / 100.0)
}

func clamp(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}