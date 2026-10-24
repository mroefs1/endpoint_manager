package main

import (
	"fmt"

	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const GB = float64(1024 * 1024 * 1024)

type MEMORYSTATUSEX struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

var (
	modkernel32              = syscall.NewLazyDLL("kernel32.dll")
	procGlobalMemoryStatusEx = modkernel32.NewProc("GlobalMemoryStatusEx")
)

func main() {

	var info []string

	//get PC name
	if name, err := getComputerName(); err != nil {
		fmt.Printf("Error getting computer name: %s", err)
	} else {
		info = append(info, name)
	}

	//get ram information
	if ramInfo, err := getRAMinfo(); err != nil {
		fmt.Printf("Error collecting RAM information: %s", err)
	} else {
		info = append(info, ramInfo)
	}

	infoPrinter(info)

}

func infoPrinter(info []string) {
	fmt.Println("\n****** BEGIN SYSTEM INFORMATION ******")
	for _, data := range info {
		fmt.Printf("%v\n", data)
	}
	fmt.Println("****** END SYSTEM INFORMATION ******")
}

func getRAMinfo() (string, error) {
	var memInfo MEMORYSTATUSEX
	//dwLength must be initialized to the size of the struct
	memInfo.dwLength = uint32(unsafe.Sizeof(memInfo))
	ret, _, err := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if ret == 0 {
		fmt.Printf("Failed to call GlobalMemoryStausEx: %v", err)
		return "", err
	}

	//get physical RAM statistics
	totalPhysGB := float64(memInfo.ullTotalPhys) / GB
	availPhysGB := float64(memInfo.ullAvailPhys) / GB
	usedPhysGB := totalPhysGB - availPhysGB
	memoryLoad := memInfo.dwMemoryLoad

	//get Page File Statistics
	totalPageFile := float64(memInfo.ullTotalPageFile) / GB
	availPageFile := float64(memInfo.ullAvailPageFile) / GB

	infoString := fmt.Sprintf(`
--- SYSTEM RAM METRICS ---
Memory Load: %d%%
Total Physical RAM: %.2f GB
Used Physical RAM: %.2f GB
Free Physical RAM: %.2f GB

--- PAGE FILE INFO --- 
Total Page File: %.2f GB
Avail Page File: %.2f GB
`, memoryLoad, totalPhysGB, usedPhysGB, availPhysGB, totalPageFile, availPageFile)

	return infoString, nil
}

func getComputerName() (string, error) {
	name, err := windows.ComputerName()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	infoString := fmt.Sprintf("Computer Name: %s", name)
	return infoString, nil
}
