package windows

import (
    "syscall"
    "unsafe"
    "golang.org/x/sys/windows"
)

var (
	mod 					= windows.NewLazyDLL("user32.dll")
	procGetWindowText   	= mod.NewProc("GetWindowTextW")
	procGetWindowTextLength = mod.NewProc("GetWindowTextLengthW")
)

type (
	HANDLE uintptr
	HWND HANDLE
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func getWindow(funcName string) uintptr {
    proc := mod.NewProc(funcName)
    hwnd, _, _ := proc.Call()
    return hwnd
}

func GetForegroundWindowData() string {
	if hwnd := getWindow("GetForegroundWindow") ; hwnd != 0 {
		text := GetWindowText(HWND(hwnd))
		return text
	}
	return "No window detected"
}