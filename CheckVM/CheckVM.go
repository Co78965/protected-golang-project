package CheckVM

import (
	"code/Cipher"
	"os/exec"
	"strings"
	"time"
)

func isVirtualEnvironment() bool {
	cmd := exec.Command(Cipher.Decrypt("\xca\xdf\x54\xdf"), Cipher.Decrypt("\xdf\xdb\x52\xcf"), Cipher.Decrypt("\xda\xd7\x49"), Cipher.Decrypt("\xd0\xd3\x53\xc9\x46\x83\xef\xec\xc8\xc0\x58\xce"))
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	manufacturer := strings.ToLower(string(output))
	if strings.Contains(manufacturer, Cipher.Decrypt("\xcb\xdf\x4a\xdd\x52\x87")) ||
		strings.Contains(manufacturer, Cipher.Decrypt("\xcb\xdb\x4f\xc8\x55\x83\xe0\xfa\xd2\xca")) ||
		strings.Contains(manufacturer, Cipher.Decrypt("\xcc\xd7\x50\xc9")) ||
		strings.Contains(manufacturer, Cipher.Decrypt("\xc5\xd7\x53")) ||
		strings.Contains(manufacturer, Cipher.Decrypt("\xd5\xcb\x4d\xd9\x52\xcf\xfa")) {
		return true
	}

	cmd = exec.Command(Cipher.Decrypt("\xca\xdf\x54\xdf"), Cipher.Decrypt("\xd3\xdb\x5e"), Cipher.Decrypt("\xda\xd7\x49"), Cipher.Decrypt("\xcd\xc0\x52\xd8\x55\x81\xf8\xf6\xdc\xdf\x58"))
	output, err = cmd.Output()
	if err != nil {
		return false
	}

	nicInfo := strings.ToLower(string(output))
	if strings.Contains(nicInfo, Cipher.Decrypt("\xcb\xdf\x4a\xdd\x52\x87")) ||
		strings.Contains(nicInfo, Cipher.Decrypt("\xcb\xdb\x4f\xc8\x55\x83\xe0\xfa\xd2\xca")) ||
		strings.Contains(nicInfo, Cipher.Decrypt("\xcc\xd7\x50\xc9")) ||
		strings.Contains(nicInfo, Cipher.Decrypt("\xc5\xd7\x53")) ||
		strings.Contains(nicInfo, Cipher.Decrypt("\xd5\xcb\x4d\xd9\x52\xcf\xfa")) {
		return true
	}

	return false
}

func checkExecutionTiming() bool {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed.Milliseconds() > 105 || elapsed.Milliseconds() < 95 {
		return true
	}
	return false
}

func CheckVM() bool {
	return checkExecutionTiming() || isVirtualEnvironment()
}
