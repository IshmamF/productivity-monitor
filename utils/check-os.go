package utils

import (
	"runtime"
)

func Get_OS() string {
	return runtime.GOOS
}
