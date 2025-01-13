package CheckDebug

import (
	"code/Cipher"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	a                              = Cipher.Decrypt
	kernel32DLL                    = syscall.NewLazyDLL(a("\xd6\xd7\x4f\xd2\x45\x8e\xbf\xaa\x93\xd6\x51\xd0"))
	isDebugger                     = kernel32DLL.NewProc(a("\xf4\xc1\x79\xd9\x42\x97\xeb\xff\xd8\xc0\x6d\xce\x45\x91\xe9\xf6\xc9"))
	procCheckRemoteDebuggerPresent = kernel32DLL.NewProc(a("\xfe\xda\x58\xdf\x4b\xb0\xe9\xf5\xd2\xc6\x58\xf8\x45\x80\xf9\xff\xda\xd7\x4f\xec\x52\x87\xff\xfd\xd3\xc6"))
)

func detectDebuggerByProcesses() bool {
	aa := Cipher.Decrypt
	out, err := exec.Command(aa("\xc9\xd3\x4e\xd7\x4c\x8b\xff\xec")).Output()
	if err != nil {
		return false
	}

	processList := string(out)
	if strings.Contains(processList, aa("\xda\xd6\x5f")) || strings.Contains(processList, aa("\xd4\xd6\x5c")) || strings.Contains(processList, aa("\xc5\x84\x09\xd8\x42\x85")) {
		return true
	}

	return false
}

func checkRemoteDebuggerPresent() bool {
	var debuggerPresent byte
	handle := syscall.Handle(0xFFFFFFFFFFFFFFFF)

	ppp := procCheckRemoteDebuggerPresent.Call
	_, _, _ = ppp(uintptr(handle), uintptr(unsafe.Pointer(&debuggerPresent)))

	return debuggerPresent != 0
}

func isDebuggerPresent1() bool {
	ok := isDebugger.Call
	flag, _, _ := ok()
	return flag != 0
}

func isDebuggerPresent() bool {
	e := isDebuggerPresent1
	return e()
}

func detectDebuggerByTiming() bool {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
	}

	elapsed := time.Since(start)

	return elapsed > (time.Millisecond)
}

func CheckDebug() bool {
	a := detectDebuggerByTiming
	b := isDebuggerPresent
	c := checkRemoteDebuggerPresent
	d := detectDebuggerByProcesses
	return (a() || b() || c() || d())
}
