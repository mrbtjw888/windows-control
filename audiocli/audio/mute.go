package audio

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

func SetMute(target string, mute bool) {
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

		vol.SetMute(mute, nil)
		fmt.Println("Mute:", name, mute)
	}
}