package v1

// DriveType is the kind of volume, as reported by GetDriveType.
// The collector maps Win32 DRIVE_* codes to these values.
type DriveType string

const (
	// DriveFixed is a local fixed disk: internal HDD, SSD or NVMe.
	DriveFixed DriveType = "fixed"
	// DriveRemovable is removable media such as a USB stick or SD card.
	// Some USB hard drives report as fixed instead.
	DriveRemovable DriveType = "removable"
	// DriveNetwork is a mapped network share. Mapped drives are
	// per logon session, so the agent (running as SYSTEM) normally
	// does not see the drives users have mapped.
	DriveNetwork DriveType = "network"
	// DriveCDROM is an optical drive.
	DriveCDROM DriveType = "cdrom"
	// DriveRAMDisk is a RAM disk.
	DriveRAMDisk DriveType = "ramdisk"
	// DriveUnknown means Windows returned a type the agent does not
	// recognize.
	DriveUnknown DriveType = "unknown"
)

// Disk is one logical volume with a drive letter.
// Volumes without a drive letter (recovery and EFI partitions,
// folder mount points) are not included. Drives with no media
// inserted are skipped.
type Disk struct {
	// Root is the drive root with its trailing backslash, e.g. "C:\".
	Root string `json:"root"`
	// Type is the kind of volume.
	Type DriveType `json:"type"`
	// Label is the volume label from GetVolumeInformation.
	// Empty if the volume has no label.
	Label string `json:"label"`
	// FileSystem is the file system name from GetVolumeInformation,
	// e.g. "NTFS", "ReFS", "FAT32" or "exFAT".
	FileSystem string `json:"fileSystem"`
	// VolumeSerial is the volume serial number from
	// GetVolumeInformation, formatted as uppercase hex "XXXX-XXXX"
	// to match the output of the vol command. It changes when the
	// volume is reformatted, so it can detect a reformat but is not
	// a hardware identifier.
	VolumeSerial string `json:"volumeSerial"`
	// TotalBytes is the volume's total size in bytes, from
	// GetDiskFreeSpaceEx.
	TotalBytes uint64 `json:"totalBytes"`
	// FreeBytes is the total free space on the volume in bytes, from
	// GetDiskFreeSpaceEx (lpTotalNumberOfFreeBytes). This ignores
	// per-user disk quotas. 0 means the volume is full.
	FreeBytes uint64 `json:"freeBytes"`
}
