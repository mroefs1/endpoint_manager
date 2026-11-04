package v1

// Scope is where a program is installed, based on which registry
// hive its Uninstall entry was found in.
type Scope string

const (
	// SoftwareScopeMachine is a machine-wide install, found under
	// HKLM and available to all users.
	SoftwareScopeMachine Scope = "machine"
	// SoftwareScopeUser is a per-user install, found under a user's
	// hive in HKEY_USERS and available only to that user.
	SoftwareScopeUser Scope = "user"
)

// Software is one installed program, as listed in Apps & features
// (Programs and Features). Read from the Uninstall registry keys:
// HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall in both
// the 64-bit and 32-bit (WOW6432Node) registry views for
// machine-wide installs, and the same key under each loaded user
// hive in HKEY_USERS for per-user installs. Entries with no
// DisplayName, and system components and updates hidden from
// Apps & features, are not included.
// Microsoft Store (MSIX/AppX) apps are not included, though Apps & features shows them.
type Software struct {
	// Name is the program's display name, e.g. "Google Chrome", from
	// the DisplayName value.
	Name string `json:"name"`
	// Version is the program's version string, e.g. "130.0.6723.92",
	// from DisplayVersion. The format is up to the vendor, so it is
	// not always comparable between programs. Empty if missing.
	Version string `json:"version"`
	// Publisher is the vendor name, e.g. "Google LLC", from the
	// Publisher value. Empty if missing.
	Publisher string `json:"publisher"`
	// InstallDate is when the program was installed or last updated,
	// as "YYYY-MM-DD". The registry stores it as YYYYMMDD; the
	// collector converts it. Installers set it themselves, so it can
	// be missing or wrong. Empty if missing or unparseable.
	InstallDate string `json:"installDate"`
	// EstimatedSizeBytes is the installer's estimate of the program's
	// size on disk, from EstimatedSize (which reports KB; the
	// collector converts to bytes). An estimate only, not a
	// measurement. 0 if missing.
	EstimatedSizeBytes uint64 `json:"estimatedSizeBytes"`
	// Architecture is the registry view the entry was found in:
	// amd64 or arm64 for the native view, x86 for WOW6432Node. It
	// reflects where the installer registered the program, which
	// usually but not always matches the program's binaries.
	Architecture Architecture `json:"architecture"`
	// Scope is where the program is installed: "machine" for entries
	// under HKLM, available to all users, or "user" for entries under
	// a user's hive, available only to that user.
	Scope Scope `json:"scope"`
	// UserSID identifies the user whose hive the entry was found in.
	// Set only when Scope is "user". Only hives loaded at collection
	// time are read, so per-user installs of users who are not
	// logged on are usually missing.
	UserSID string `json:"userSid,omitempty"`
}
