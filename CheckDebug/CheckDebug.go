package CheckDebug

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32DLL                    = syscall.NewLazyDLL("kernel32.dll")
	isDebugger                     = kernel32DLL.NewProc("IsDebuggerPresent")
	procCheckRemoteDebuggerPresent = kernel32DLL.NewProc("CheckRemoteDebuggerPresent")
)

func detectDebuggerByProcesses() bool {
	// Используем команду 'tasklist' для Windows или 'ps' для UNIX-подобных систем
	out, err := exec.Command("tasklist").Output() // Для UNIX заменить на "ps aux"
	if err != nil {
		fmt.Println("Ошибка при выполнении команды:", err)
		return false
	}

	processList := string(out)
	// Проверяем наличие известных процессов дебаггеров
	if strings.Contains(processList, "gdb") || strings.Contains(processList, "ida") || strings.Contains(processList, "x64dbg") {
		return true
	}

	return false
}

// if true --> debug detected
func ntQueryInformationProcess() bool {
	var debuggerPresent byte
	handle := syscall.Handle(0xFFFFFFFFFFFFFFFF) // Текущий процесс
	_, _, _ = procCheckRemoteDebuggerPresent.Call(uintptr(handle), uintptr(unsafe.Pointer(&debuggerPresent)))

	return debuggerPresent != 0
}

// if true --> debug detected
func isDebuggerPresent1() bool {
	flag, _, _ := isDebugger.Call()
	return flag != 0
}

func isDebuggerPresent() bool {
	return isDebuggerPresent1()
}

// Выполняется при дебаге примерно за >3мс
func detectDebuggerByTiming() bool {
	start := time.Now()

	// Простой код, выполнение которого не должно занимать много времени
	for i := 0; i < 1000000; i++ {
	}

	elapsed := time.Since(start)
	fmt.Printf("Время выполнения: %s\n", elapsed)

	// Предположим, что если выполнение заняло больше 100 миллисекунд, это может быть из-за дебаггера
	return elapsed > (time.Millisecond)
}

func CheckDebug() bool {
	return (detectDebuggerByTiming() || isDebuggerPresent() || ntQueryInformationProcess() || detectDebuggerByProcesses())
}
