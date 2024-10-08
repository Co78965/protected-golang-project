package CheckVM

import (
	"os"
	"strings"
	"time"
)

func checkDevice(deviceName string) bool {
	file, err := os.Open("/proc/scsi/scsi") // файл с информацией о подключенных устройствах в Linux
	if err != nil {
		return false
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	content := string(buf[:n])

	return strings.Contains(content, deviceName)
}

func CheckVMware() bool {
	devices := []string{"VMware", "VBOX"}

	for _, device := range devices {
		if checkDevice(device) {
			return true
		}
	}
	return false
}

func CheckExecutionTiming() bool {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(start)

	// Если время отличается слишком сильно от 100ms, это может указывать на виртуальную среду
	if elapsed.Milliseconds() > 105 || elapsed.Milliseconds() < 95 {
		return true
	}
	return false
}
