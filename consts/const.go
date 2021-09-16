package consts

import "runtime"

var (
	Which string
)

func init() {
	switch runtime.GOOS {
	case "linux":
		Which = "which"
	case "windows":
		Which = "where"
	default:
		Which = "which"
	}
}
