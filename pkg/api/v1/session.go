package v1

import "time"

// SessionState is a session's connection state, from the State field
// reported by WTSEnumerateSessionsEx. The collector maps Win32
// WTS_CONNECTSTATE_CLASS values to these values.
type SessionState string

const (
	// SessionStateActive means a user is logged on and connected,
	// either at the console or over RDP.
	SessionStateActive SessionState = "active"
	// SessionStateDisconnected means a user is logged on but not
	// connected, e.g. an RDP client was closed without signing out
	// or the user switched to another account. Their programs keep
	// running.
	SessionStateDisconnected SessionState = "disconnected"
	// SessionStateConnected means a client is connected but no user
	// has logged on yet.
	SessionStateConnected SessionState = "connected"
	// SessionStateConnectQuery means the session is in the process of
	// connecting to a client.
	SessionStateConnectQuery SessionState = "connectQuery"
	// SessionStateShadow means the session is remotely controlling
	// (shadowing) another session.
	SessionStateShadow SessionState = "shadow"
	// SessionStateIdle means the session is waiting for a client to
	// connect.
	SessionStateIdle SessionState = "idle"
	// SessionStateReset means the session is being reset.
	SessionStateReset SessionState = "reset"
	// SessionStateDown means the session is down due to an error.
	SessionStateDown SessionState = "down"
	// SessionStateInit means the session is initializing.
	SessionStateInit SessionState = "init"
	// SessionStateUnknown means Windows reported a state the agent
	// does not recognize.
	SessionStateUnknown SessionState = "unknown"
)

// Session is one Windows logon session, from WTSEnumerateSessionsEx.
// Listener sessions (e.g. "RDP-Tcp") are not included. Session 0 is
// where services run and normally has no user.
type Session struct {
	// ID is the session ID. 0 is the services session; user sessions
	// start at 1. IDs are reused after a session ends, so use ID
	// together with LogonTime to tell sessions apart across reports.
	// Session 0 is intentionally omitted, similarly to process.go
	ID uint32 `json:"id"`
	// UserName is the logged-on user's account name, without the
	// domain, e.g. "jsmith". Empty if no user is logged on.
	UserName string `json:"userName"`
	// Domain is the user's domain: the NetBIOS domain name for domain
	// accounts, the computer name for local accounts, or "AzureAD"
	// for Entra ID accounts. Empty if no user is logged on.
	Domain string `json:"domain"`
	// State is the session's connection state.
	State SessionState `json:"state"`
	// StationName is the WinStation name, e.g. "Console" for the
	// physical console or "RDP-Tcp#3" for a remote session. Empty for
	// disconnected sessions.
	StationName string `json:"stationName"`
	// ClientName is the computer name of the remote client, from
	// WTSQuerySessionInformation (WTSClientName). Empty for console
	// sessions.
	ClientName string `json:"clientName"`
	// LogonTime is when the user logged on, in UTC, from
	// WTSQuerySessionInformation (WTSSessionInfo). Omitted if no user
	// is logged on.
	LogonTime time.Time `json:"logonTime,omitzero"`
}
