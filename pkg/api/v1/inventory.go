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
	CollectorOS              CollectorName = "os"
	CollectorHardware        CollectorName = "hardware"
	CollectorDisks           CollectorName = "disks"
	CollectorNetworkAdapters CollectorName = "networkAdapters"
	CollectorSessions        CollectorName = "sessions"
	CollectorServices        CollectorName = "services"
	CollectorProcesses       CollectorName = "processes"
	CollectorSoftware        CollectorName = "software"
	CollectorHealth          CollectorName = "health"
)

// CollectorResult records one collector run. Error is empty on
// success. A section whose collector failed is null in the JSON.
type CollectorResult struct {
	Name  CollectorName `json:"name"`
	Error string        `json:"error,omitempty"`
	// How long the collector took, measured in milliseconds
	DurationMs int64 `json:"durationMs"`
}

// Inventory is the full snapshot sent to the server.
// Keep JSON friendly so Go and Flutter can share it.
// MachineID is the identity key, times are UTC.
// Collectors is the source of truth for whether each section was successfully collected.
type Inventory struct {
	// From SchemaVersion constant
	SchemaVersion int `json:"schemaVersion"`
	// Agent's semantic version
	AgentVersion string `json:"agentVersion"`
	// From HKLM\SOFTWARE\Microsoft\Cryptography\MachineGuid.
	// Read from 64-bit registry view and normalized to lowercase without braces.
	// It is the identity key for the endpoint.
	MachineID string `json:"machineId"`
	// The DNS hostname, for display only. Never use as an identifier, hostnames change.
	Hostname string `json:"hostname"`
	// When the collection started, in UTC
	CollectedAt     time.Time        `json:"collectedAt"`
	OS              *OSInfo          `json:"os"`
	Hardware        *HardwareInfo    `json:"hardware"`
	Disks           []Disk           `json:"disks"`
	NetworkAdapters []NetworkAdapter `json:"networkAdapters"`
	Sessions        []Session        `json:"sessions"`
	Services        []Service        `json:"services"`
	Processes       []Process        `json:"processes"`
	Software        []Software       `json:"software"`
	Health          *HealthInfo      `json:"health"`

	// Collectors has one entry per collector that ran, with its error (if any) and duration
	Collectors []CollectorResult `json:"collectors"`
}
