package audio

import (
	"strings"

	"golang.org/x/sys/windows"
)

func getProcessName(pid uint32) string {
	handle, err := windows.OpenProcess(
		windows.PROCESS_QUERY_LIMITED_INFORMATION,
		false,
		pid,
	)
	if err != nil {
		return ""
	}
	defer windows.CloseHandle(handle)

	buf := make([]uint16, 260)
	size := uint32(len(buf))

	err = windows.QueryFullProcessImageName(
		handle,
		0,
		&buf[0],
		&size,
	)
	if err != nil {
		return ""
	}

	full := windows.UTF16ToString(buf[:size])
	return extractExe(full)
}

func extractExe(path string) string {
	path = strings.ReplaceAll(path, "/", "\\")
	parts := strings.Split(path, "\\")
	return parts[len(parts)-1]
}