package v1

// RebootReason is one reason Windows is waiting for a restart. Each
// value matches one registry location the collector checks.
type RebootReason string

const (
	// RebootReasonWindowsUpdate means Windows Update installed updates
	// that need a restart: the RebootRequired key exists under
	// HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\WindowsUpdate\Auto Update.
	RebootReasonWindowsUpdate RebootReason = "windowsUpdate"
	// RebootReasonCBS means Component Based Servicing (the servicing
	// stack behind Windows features, roles and cumulative updates) has
	// work to finish at restart: the RebootPending key exists under
	// HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Component Based Servicing.
	RebootReasonCBS RebootReason = "cbs"
	// RebootReasonFileRename means files are queued to be replaced or
	// deleted at restart: the PendingFileRenameOperations value under
	// HKLM\SYSTEM\CurrentControlSet\Control\Session Manager is set.
	// Many installers and antivirus products set it, so on its own it
	// is a weak signal.
	RebootReasonFileRename RebootReason = "fileRename"
	// RebootReasonComputerRename means the computer was renamed and the
	// new name takes effect at restart: ComputerName under
	// HKLM\SYSTEM\CurrentControlSet\Control\ComputerName\ComputerName
	// differs from the one under ...\ComputerName\ActiveComputerName.
	RebootReasonComputerRename RebootReason = "computerRename"
)

// SecureBootState is whether UEFI Secure Boot is enforcing, from the
// UEFISecureBootEnabled value under
// HKLM\SYSTEM\CurrentControlSet\Control\SecureBoot\State.
type SecureBootState string

const (
	// SecureBootStateOn means Secure Boot is enabled and enforcing.
	SecureBootStateOn SecureBootState = "on"
	// SecureBootStateOff means the firmware supports Secure Boot but it
	// is turned off.
	SecureBootStateOff SecureBootState = "off"
	// SecureBootStateUnsupported means the machine cannot use Secure
	// Boot, typically because it boots in legacy BIOS mode.
	SecureBootStateUnsupported SecureBootState = "unsupported"
	// SecureBootStateUnknown means the state could not be read.
	SecureBootStateUnknown SecureBootState = "unknown"
)

// FirmwareType is the firmware interface Windows booted through, from
// GetFirmwareType. The same machine can report either value depending
// on how the firmware is configured to boot.
type FirmwareType string

const (
	// FirmwareTypeUEFI means Windows booted in UEFI mode. Required for
	// Secure Boot and for Windows 11.
	FirmwareTypeUEFI FirmwareType = "uefi"
	// FirmwareTypeBIOS means Windows booted in legacy BIOS mode,
	// including UEFI firmware running its compatibility module (CSM).
	FirmwareTypeBIOS FirmwareType = "bios"
	// FirmwareTypeUnknown means Windows returned a type the agent does
	// not recognize.
	FirmwareTypeUnknown FirmwareType = "unknown"
)

// Firewall is whether Windows Defender Firewall is enabled for each
// network profile, from INetFwPolicy2. This is the effective state,
// including Group Policy. A third-party firewall that replaces
// Windows Defender Firewall shows here as disabled.
type Firewall struct {
	// Domain is the profile used on networks where the machine can
	// reach its Active Directory domain controller.
	Domain bool `json:"domain"`
	// Private is the profile for networks the user marked as private,
	// e.g. a home network. The registry calls it StandardProfile.
	Private bool `json:"private"`
	// Public is the profile for untrusted networks, and the default for
	// networks Windows has not classified.
	Public bool `json:"public"`
}

// HealthInfo is a summary of the machine's security and maintenance
// state.
type HealthInfo struct {
	// RebootReasons lists every reason a restart is pending. Empty if
	// no restart is pending.
	RebootReasons []RebootReason `json:"rebootReasons"`
	// Firewall is the firewall state for each network profile.
	Firewall Firewall `json:"firewallProfiles"`
	// UACEnabled is whether User Account Control is on, from the
	// EnableLUA value under
	// HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System.
	// Setting the UAC slider to "Never notify" does not turn it off,
	// so this stays true. Changes take effect only after a restart,
	// so this can reflect a pending setting.
	UACEnabled bool `json:"uacEnabled"`
	// SecureBoot is whether Secure Boot is enforcing.
	SecureBoot SecureBootState `json:"secureBoot"`
	// FirmwareType is the firmware interface Windows booted through.
	FirmwareType FirmwareType `json:"firmwareType"`
}
