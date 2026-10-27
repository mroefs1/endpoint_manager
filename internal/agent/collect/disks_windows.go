package collect

import (
	apiv1 "github.com/mroefs1/endpoint_manager/pkg/api/v1"
	"golang.org/x/sys/windows"
)

var driveTypes = map[uint32]apiv1.DriveType{
	windows.DRIVE_FIXED:     apiv1.DriveFixed,
	windows.DRIVE_REMOVABLE: apiv1.DriveRemovable,
	windows.DRIVE_REMOTE:    apiv1.DriveNetwork,
	windows.DRIVE_CDROM:     apiv1.DriveCDROM,
	windows.DRIVE_RAMDISK:   apiv1.DriveRAMDisk,
	windows.DRIVE_UNKNOWN:   apiv1.DriveUnknown,
}

// Disks returns every lettered volume with media present.
// Drives that disappear or have no media during collection are skipped, not reported as errors.
func Disks() ([]apiv1.Disk, error) {
	// REMINDER FOR IMPLEMENTATION: DRIVE_NO_ROOT_DIR isn’t in the map, which is correct,
	// so check for it and continue before the lookup. Keep the two-value lookup
	// with the DriveUnknown fallback for anything else.

	disks := []apiv1.Disk{}
	return disks, nil
}
