package v1

// HardwareInfo describes the physical (or virtual) machine.
// Manufacturer, Model, SerialNumber and SMBIOSUUID come from the
// SMBIOS System Information structure (type 1), read with
// GetSystemFirmwareTable. Firmware fills these in, so on custom-built
// PCs they may hold placeholders such as "To Be Filled By O.E.M." or
// "Default string". On virtual machines they describe the VM.
type HardwareInfo struct {
	// Manufacturer is the system maker, e.g. "Dell Inc.", "LENOVO",
	// or "VMware, Inc." on a VMware VM.
	Manufacturer string `json:"manufacturer"`
	// Model is the SMBIOS Product Name, e.g. "Latitude 7440". Some
	// vendors put a machine type code here instead of the marketing
	// name (Lenovo uses codes like "21HD003BUS").
	Model string `json:"model"`
	// SerialNumber is the system serial number (the service tag on
	// Dell machines). Can be empty or a placeholder on custom builds
	// and some VMs.
	SerialNumber string `json:"serialNumber"`
	// SMBIOSUUID is the system UUID, formatted as uppercase
	// 8-4-4-4-12 hex to match what Windows tools display. Informational
	// only, never an identifier: cloned VMs and some vendors' boards
	// share UUIDs. Empty if firmware reports it as not set.
	SMBIOSUUID string `json:"smbiosUuid"`
	// BIOS is the system firmware.
	BIOS BIOSInfo `json:"bios"`
	// CPU summarizes the installed processors.
	CPU CPUInfo `json:"cpu"`
	// Memory is the system's physical memory.
	Memory MemoryInfo `json:"memory"`
}

// BIOSInfo is the system firmware (BIOS or UEFI), from the SMBIOS
// BIOS Information structure (type 0).
type BIOSInfo struct {
	// Vendor is the firmware maker, e.g. "Dell Inc." or
	// "American Megatrends Inc.".
	Vendor string `json:"vendor"`
	// Version is the firmware version string, e.g. "1.14.0". The
	// format is vendor-specific, so compare versions only between
	// machines of the same model.
	Version string `json:"version"`
	// ReleaseDate is the firmware release date as "YYYY-MM-DD".
	// SMBIOS stores it as MM/DD/YYYY; the collector converts it.
	// Empty if missing or unparseable.
	ReleaseDate string `json:"releaseDate"`
}

// CPUInfo summarizes the installed processors. Multi-socket machines
// are assumed to have identical processors in each socket.
type CPUInfo struct {
	// Name is the processor model, e.g.
	// "13th Gen Intel(R) Core(TM) i7-1365U", from ProcessorNameString
	// under HKLM\HARDWARE\DESCRIPTION\System\CentralProcessor\0.
	// Leading and trailing spaces are trimmed.
	Name string `json:"name"`
	// Sockets is the number of physical processor packages, from
	// GetLogicalProcessorInformationEx.
	Sockets uint32 `json:"sockets"`
	// Cores is the number of physical cores across all sockets, from
	// GetLogicalProcessorInformationEx.
	Cores uint32 `json:"cores"`
	// LogicalProcessors is the number of hardware threads across all
	// processor groups, from GetActiveProcessorCount. Higher than
	// Cores when hyperthreading or SMT is enabled.
	LogicalProcessors uint32 `json:"logicalProcessors"`
	// BaseClockMHz is the processor's base clock speed in MHz, from
	// the ~MHz value under the same CentralProcessor\0 key. Not the
	// boost or current speed.
	BaseClockMHz uint32 `json:"baseClockMhz"`
}

// MemoryInfo describes physical memory, in bytes.
type MemoryInfo struct {
	// InstalledBytes is the physically installed memory, from
	// GetPhysicallyInstalledSystemMemory (which reports KB; the
	// collector converts to bytes). Matches the machine's spec sheet.
	InstalledBytes uint64 `json:"installedBytes"`
	// TotalBytes is the memory usable by Windows, from
	// GlobalMemoryStatusEx. Lower than InstalledBytes by the amount
	// reserved for hardware and integrated graphics.
	TotalBytes uint64 `json:"totalBytes"`
	// AvailableBytes is memory available to programs at collection
	// time (free plus standby cache), from GlobalMemoryStatusEx. A
	// point-in-time reading, same as Task Manager's "Available".
	AvailableBytes uint64 `json:"availableBytes"`
}
