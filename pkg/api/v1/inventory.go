package v1

import (
	"time"
)

// SchemaVersion is the version of the Inventory JSON contract.
// Bump only for breaking changes (removing or renaming fields).
// Adding fields does not require a bump.
const SchemaVersion = 1

// CollectorName identifies one collector. Each value matches the
// JSON key of the Inventory section that collector fills.
type CollectorName string

const (
	// CollectorOS fills Inventory.OS.
	CollectorOS              CollectorName = "os"
	// CollectorHardware fills Inventory.Hardware.
	CollectorHardware        CollectorName = "hardware"
	// CollectorDisks fills Inventory.Disks.
	CollectorDisks           CollectorName = "disks"
	// CollectorNetworkAdapters fills Inventory.NetworkAdapters.
	CollectorNetworkAdapters CollectorName = "networkAdapters"
	// CollectorSessions fills Inventory.Sessions.
	CollectorSessions        CollectorName = "sessions"
	// CollectorServices fills Inventory.Services.
	CollectorServices        CollectorName = "services"
	// CollectorProcesses fills Inventory.Processes.
	CollectorProcesses       CollectorName = "processes"
	// CollectorSoftware fills Inventory.Software.
	CollectorSoftware        CollectorName = "software"
	// CollectorHealth fills Inventory.Health.
	CollectorHealth          CollectorName = "health"
)

// CollectorResult records one collector run. Error is empty on
// success. A section whose collector failed is null in the JSON.
type CollectorResult struct {
	// Name is the collector that ran. It matches the JSON key of the
	// Inventory section the collector fills.
	Name  CollectorName `json:"name"`
	// Error is the error message if the collector failed, e.g.
	// "read SMBIOS table: ...". Omitted on success.
	Error string        `json:"error,omitempty"`
	// DurationMs is how long the collector took to run, in
	// milliseconds. Set whether the collector succeeded or failed.
	DurationMs int64 `json:"durationMs"`
}

// Inventory is the full snapshot of one machine, sent to the server.
// Keep it JSON friendly so Go and Flutter can share it. MachineID is
// the identity key, and all times are UTC.
// Each section (OS through Health) is filled by one collector and is
// null if that collector failed. Collectors is the source of truth
// for whether each section was collected successfully.
type Inventory struct {
	// SchemaVersion is the version of this JSON contract, set from the
	// SchemaVersion constant. The server uses it to tell which contract
	// a report follows.
	SchemaVersion int `json:"schemaVersion"`
	// AgentVersion is the agent's semantic version, e.g. "0.1.0".
	AgentVersion string `json:"agentVersion"`
	// MachineID is the MachineGuid value under
	// HKLM\SOFTWARE\Microsoft\Cryptography, read from the 64-bit
	// registry view and normalized to lowercase without braces. It is
	// the identity key for the endpoint.
	MachineID string `json:"machineId"`
	// Hostname is the machine's DNS hostname, e.g. "pc01". Display
	// only: hostnames change, so never use it as an identifier.
	Hostname string `json:"hostname"`
	// CollectedAt is when the collection started, in UTC.
	CollectedAt     time.Time        `json:"collectedAt"`
	// OS is the installed operating system. Null if the os collector
	// failed.
	OS              *OSInfo          `json:"os"`
	// Hardware is the physical (or virtual) machine. Null if the
	// hardware collector failed.
	Hardware        *HardwareInfo    `json:"hardware"`
	// Disks are the logical volumes with a drive letter. Null if the
	// disks collector failed.
	Disks           []Disk           `json:"disks"`
	// NetworkAdapters are the network interfaces, excluding loopback.
	// Null if the networkAdapters collector failed.
	NetworkAdapters []NetworkAdapter `json:"networkAdapters"`
	// Sessions are the logon sessions with a logged-on user. Null if
	// the sessions collector failed.
	Sessions        []Session        `json:"sessions"`
	// Services are the installed Win32 services. Null if the services
	// collector failed.
	Services        []Service        `json:"services"`
	// Processes are the processes running at collection time. Null if
	// the processes collector failed.
	Processes       []Process        `json:"processes"`
	// Software is the installed programs, as listed in Apps & features.
	// Null if the software collector failed.
	Software        []Software       `json:"software"`
	// Health is the machine's security and maintenance state. Null if
	// the health collector failed.
	Health          *HealthInfo      `json:"health"`

	// Collectors has one entry per collector that ran, with its error
	// (if any) and duration. Use it, not the sections, to tell whether
	// each section was collected successfully.
	Collectors []CollectorResult `json:"collectors"`
}
