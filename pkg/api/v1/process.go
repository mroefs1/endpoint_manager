package v1

import "time"

// Process is one running process, from a CreateToolhelp32Snapshot
// process snapshot. The list is a point-in-time reading: processes
// that start or exit during collection may be missing.
type Process struct {
	// PID is the process ID. Windows reuses PIDs after a process
	// exits, so use PID together with StartTime to tell processes
	// apart across reports.
	// The System Idle Process (PID 0) is not included; PID 4 is the System process.
	PID uint32 `json:"pid"`
	// ParentPID is the PID of the process that created this one, from
	// th32ParentProcessID. Windows does not update it when the parent
	// exits, so the parent may be gone or the PID may now belong to an
	// unrelated process. Treat it as the parent only if that process
	// has an earlier StartTime.
	ParentPID uint32 `json:"parentPid"`
	// Name is the executable file name without its path, e.g.
	// "svchost.exe", from szExeFile.
	Name string `json:"name"`
	// ImagePath is the full path of the executable, e.g.
	// "C:\Windows\System32\svchost.exe", from
	// QueryFullProcessImageName. Empty for the
	// the System process and processes the agent cannot open.
	ImagePath string `json:"imagePath"`
	// UserName is the account name the process runs as, without the
	// domain, e.g. "jsmith" or "SYSTEM". Read from the process token
	// and resolved with LookupAccountSid. Empty if the token cannot
	// be read.
	UserName string `json:"userName"`
	// Domain is the account's domain, in the same format as
	// Session.Domain: the NetBIOS domain name for domain accounts,
	// the computer name for local accounts, "AzureAD" for Entra ID
	// accounts, or "NT AUTHORITY" for built-in accounts such as
	// SYSTEM. Empty if the token cannot be read.
	Domain string `json:"domain"`
	// SessionID is the session the process runs in, from
	// ProcessIdToSessionId. 0 is the services session; match it
	// against Session.ID to find the user session a process belongs
	// to. Session 0 is not in the list.
	SessionID uint32 `json:"sessionId"`
	// StartTime is when the process started, in UTC, from the creation
	// time reported by GetProcessTimes. Omitted if the process cannot
	// be opened.
	StartTime time.Time `json:"startTime,omitzero"`
}
