package v1

import "time"

// ProductType is the role of the Windows installation, from the
// wProductType field returned by RtlGetVersion.
type ProductType string

const (
	// ProductWorkstation is a client edition (Windows 10, Windows 11).
	ProductWorkstation ProductType = "workstation"
	// ProductDomainController is a server acting as an Active
	// Directory domain controller.
	ProductDomainController ProductType = "domainController"
	// ProductServer is a Windows Server edition that is not a
	// domain controller.
	ProductServer ProductType = "server"
	// ProductUnknown means Windows returned a value the agent does
	// not recognize.
	ProductUnknown ProductType = "unknown"
)

// Architecture is the native CPU architecture of the operating
// system, not of the agent binary. A 32-bit agent on 64-bit Windows
// still reports amd64.
type Architecture string

const (
	// ArchitectureAMD64 is 64-bit x86 (what Windows calls x64).
	ArchitectureAMD64 Architecture = "amd64"
	// ArchitectureARM64 is 64-bit ARM.
	ArchitectureARM64 Architecture = "arm64"
	// ArchitectureX86 is 32-bit x86.
	ArchitectureX86 Architecture = "x86"
	// ArchitectureUnknown means Windows returned a machine type the
	// agent does not recognize.
	ArchitectureUnknown Architecture = "unknown"
)

// JoinType is how the machine is joined to a directory.
type JoinType string

const (
	// JoinWorkgroup is not joined to any directory.
	JoinWorkgroup JoinType = "workgroup"
	// JoinDomain is joined to an on-premises Active Directory domain.
	JoinDomain JoinType = "domain"
	// JoinEntraID is joined to Microsoft Entra ID (formerly Azure AD)
	// only.
	JoinEntraID JoinType = "entraId"
	// JoinHybrid is joined to both an Active Directory domain and
	// Entra ID.
	JoinHybrid JoinType = "hybrid"
	// JoinUnknown means the join state could not be determined.
	JoinUnknown JoinType = "unknown"
)

// OSInfo describes the installed Windows operating system.
type OSInfo struct {
	// ProductName is the marketing name, e.g. "Windows 11 Pro" or
	// "Windows Server 2022 Standard". Read from the ProductName value
	// under HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion. The
	// registry reports "Windows 10" on Windows 11; the collector
	// corrects this using BuildNumber, so this field is accurate.
	ProductName string `json:"productName"`
	// EditionID is the edition identifier, e.g. "Professional",
	// "Enterprise" or "ServerStandard". Read from EditionID under the
	// same CurrentVersion key. Better than ProductName for filtering
	// because it does not change with marketing names.
	EditionID string `json:"editionId"`
	// DisplayVersion is the feature update release, e.g. "23H2".
	// Read from DisplayVersion under the same CurrentVersion key.
	// Empty on releases older than Windows 10 20H2, which used a
	// different value.
	DisplayVersion string `json:"displayVersion"`
	// MajorVersion is the major version from RtlGetVersion. It is 10
	// for Windows 10, Windows 11 and Windows Server 2016 and later.
	MajorVersion uint32 `json:"majorVersion"`
	// MinorVersion is the minor version from RtlGetVersion, 0 on all
	// current releases.
	MinorVersion uint32 `json:"minorVersion"`
	// BuildNumber is the OS build from RtlGetVersion, e.g. 22631.
	// Client builds 22000 and above are Windows 11. Use this, not
	// ProductName, for version comparisons.
	BuildNumber uint32 `json:"buildNumber"`
	// UBR is the update build revision: the patch level within a
	// build, the 4317 in 22631.4317. It increases with each
	// cumulative update. Read from UBR under the same CurrentVersion
	// key.
	UBR uint32 `json:"ubr"`
	// Architecture is the native OS architecture, from
	// IsWow64Process2.
	Architecture Architecture `json:"architecture"`
	// ProductType is whether this is a workstation, server or domain
	// controller.
	ProductType ProductType `json:"productType"`
	// FQDN is the fully qualified DNS name, e.g. "pc01.corp.example.com",
	// from GetComputerNameEx. On machines with no DNS suffix it equals
	// the hostname. Display only: never use it as an identifier.
	FQDN string `json:"fqdn"`
	// JoinType is how the machine is joined to a directory, combining
	// NetGetJoinInformation (Active Directory) and
	// NetGetAadJoinInformation (Entra ID).
	JoinType JoinType `json:"joinType"`
	// JoinName depends on JoinType: the AD domain name for domain and
	// hybrid, the workgroup name for workgroup, the Entra tenant
	// display name for entraId, and empty for unknown.
	JoinName string `json:"joinName"`
	// LastBootTime is when Windows last started, in UTC, computed as
	// collection time minus the system tick count. It can differ by
	// a second or so between reports, so do not compare it for exact
	// equality. With Fast Startup enabled (the client default), shutting
	// down does not reset it; only a restart does.
	LastBootTime time.Time `json:"lastBootTime"`
	// InstallDate is when the current Windows installation was set up,
	// in UTC, from InstallDate under the same CurrentVersion key.
	// Windows resets it during feature updates, so on most machines it
	// is the date of the last feature update, not the original
	// install. Omitted if the registry value is missing.
	InstallDate time.Time `json:"installDate,omitzero"`
}
