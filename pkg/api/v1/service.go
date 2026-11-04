package v1

// ServiceState is a service's current state, from the
// dwCurrentState field reported by EnumServicesStatusEx.
type ServiceState string

const (
	// ServiceStateRunning means the service is running.
	ServiceStateRunning ServiceState = "running"
	// ServiceStateStopped means the service is not running.
	ServiceStateStopped ServiceState = "stopped"
	// ServiceStateStartPending means the service is starting. A
	// service stuck in this state across several reports is hung.
	ServiceStateStartPending ServiceState = "startPending"
	// ServiceStateStopPending means the service is stopping.
	ServiceStateStopPending ServiceState = "stopPending"
	// ServiceStatePaused means the service is paused. Few services
	// support pausing.
	ServiceStatePaused ServiceState = "paused"
	// ServiceStatePausePending means the service is pausing.
	ServiceStatePausePending ServiceState = "pausePending"
	// ServiceStateContinuePending means the service is resuming
	// from a pause.
	ServiceStateContinuePending ServiceState = "continuePending"
	// ServiceStateUnknown means Windows reported a state the agent
	// does not recognize.
	ServiceStateUnknown ServiceState = "unknown"
)

// StartType is how a service is configured to start, from the
// dwStartType field reported by QueryServiceConfig.
type StartType string

const (
	// StartAutomatic means the service starts at boot.
	StartAutomatic StartType = "automatic"
	// StartAutomaticDelayed means the service starts automatically
	// shortly after boot, once other automatic services have started.
	// Windows reports these as automatic; the delayed flag comes from
	// a separate QueryServiceConfig2 call.
	StartAutomaticDelayed StartType = "automaticDelayed"
	// StartManual means the service starts only when requested by a
	// user, a program or another service.
	StartManual StartType = "manual"
	// StartDisabled means the service cannot be started.
	StartDisabled StartType = "disabled"
	// StartUnknown means Windows reported a start type the agent does
	// not recognize.
	StartUnknown StartType = "unknown"
)

// Service is one Windows service. Only Win32 services are included,
// not kernel drivers. Per-user services (e.g. "OneSyncSvc_4a2b1")
// appear once per logged-on user, each with its own suffix.
type Service struct {
	// Name is the service's key name, e.g. "wuauserv". It is stable
	// and is what sc.exe and the registry use.
	Name string `json:"name"`
	// DisplayName is the name shown in the Services console, e.g.
	// "Windows Update". It can be localized, so use Name to identify
	// a service.
	DisplayName string `json:"displayName"`
	// State is the service's current state.
	State ServiceState `json:"state"`
	// StartType is how the service is configured to start.
	StartType StartType `json:"startType"`
	// PID is the process ID the service runs in, or 0 when it is not
	// running. Several services can share one PID when they run
	// inside the same svchost.exe process.
	PID uint32 `json:"pid"`
	// Account is the account the service runs as, from
	// QueryServiceConfig, e.g. "LocalSystem",
	// "NT AUTHORITY\NetworkService", "NT SERVICE\<name>" (a virtual
	// account) or a domain account. Domain accounts are worth
	// tracking because their passwords expire.
	Account string `json:"account"`
}
