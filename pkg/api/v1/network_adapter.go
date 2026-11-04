package v1

import "net/netip"

// OperStatus is an adapter's operational state, from the OperStatus
// field of GetAdaptersAddresses. The values follow the standard
// interface states defined in RFC 2863.
type OperStatus string

const (
	// OperStatusUp means the adapter can pass traffic.
	OperStatusUp OperStatus = "up"
	// OperStatusDown means the adapter cannot pass traffic, e.g. the
	// cable is unplugged or Wi-Fi is not connected.
	OperStatusDown OperStatus = "down"
	// OperStatusTesting means the adapter is in a test mode.
	OperStatusTesting OperStatus = "testing"
	// OperStatusUnknown means Windows could not determine the state.
	OperStatusUnknown OperStatus = "unknown"
	// OperStatusDormant means the adapter is waiting for an external
	// event before it can pass traffic, e.g. a connection being
	// established.
	OperStatusDormant OperStatus = "dormant"
	// OperStatusNotPresent means a component of the adapter is
	// missing, typically a hardware or driver problem.
	OperStatusNotPresent OperStatus = "notPresent"
	// OperStatusLowerLayerDown means the adapter depends on another
	// interface that is down, e.g. a VPN adapter whose underlying
	// network connection has dropped.
	OperStatusLowerLayerDown OperStatus = "lowerLayerDown"
)

// AdapterType is the kind of network interface, mapped from the
// IfType field of GetAdaptersAddresses. Virtual adapters (Hyper-V,
// VPN clients, container networks) usually report as ethernet or
// tunnel, so this does not tell physical and virtual apart.
type AdapterType string

const (
	// AdapterEthernet is a wired Ethernet interface, or a virtual
	// adapter that presents itself as one.
	AdapterEthernet AdapterType = "ethernet"
	// AdapterWiFi is a wireless (802.11) interface.
	AdapterWiFi AdapterType = "wifi"
	// AdapterTunnel is a tunnel interface, e.g. Teredo or some VPNs.
	AdapterTunnel AdapterType = "tunnel"
	// AdapterPPP is a point-to-point interface, e.g. some VPN or
	// cellular connections.
	AdapterPPP AdapterType = "ppp"
	// AdapterOther is any other interface type Windows reports.
	AdapterOther AdapterType = "other"
)

// NetworkAdapter is one network interface, from GetAdaptersAddresses.
// Loopback interfaces are skipped.
type NetworkAdapter struct {
	// FriendlyName is the name shown in Windows network settings,
	// e.g. "Ethernet 2" or "Wi-Fi". Users can rename adapters, so it
	// is for display only.
	FriendlyName string `json:"friendlyName"`
	// Description is the adapter's driver description, e.g.
	// "Intel(R) Ethernet Connection (17) I219-LM". It can change when
	// the driver is updated.
	Description string `json:"description"`
	// MACAddress is the adapter's hardware address as lowercase
	// colon-separated hex (the net.HardwareAddr format), e.g.
	// "aa:bb:cc:dd:ee:ff". Empty for adapters without one.
	MACAddress string `json:"macAddress"`
	// Status is the adapter's operational state.
	Status OperStatus `json:"status"`
	// Type is the kind of interface.
	Type AdapterType `json:"type"`
	// LinkSpeedBps is the negotiated transmit link speed in bits per
	// second, e.g. 1000000000 for gigabit Ethernet. This is the
	// connection's rated speed, not current throughput. On Wi-Fi,
	// receive speed can differ. 0 means unknown, which is normal for
	// disconnected adapters.
	LinkSpeedBps uint64 `json:"linkSpeedBps"`
	// DHCPEnabled reports whether IPv4 addresses are assigned by DHCP.
	// IPv6 addresses are configured separately and not covered by
	// this flag.
	DHCPEnabled bool `json:"dhcpEnabled"`
	// DNSSuffix is the connection-specific DNS suffix, e.g.
	// "corp.example.com", usually assigned by DHCP. Empty if none.
	DNSSuffix string `json:"dnsSuffix"`
	// Addresses are the adapter's unicast IPv4 and IPv6 addresses,
	// each with its prefix length, e.g. "10.0.0.5/24". These are host
	// addresses, not network addresses. IPv6 zone identifiers are
	// removed.
	Addresses []netip.Prefix `json:"addresses"`
	// Gateways are the adapter's default gateways, IPv4 and IPv6.
	// Empty for adapters with no gateway.
	Gateways []netip.Addr `json:"gateways"`
	// DNSServers are the DNS servers configured on this adapter, in
	// the order Windows uses them.
	DNSServers []netip.Addr `json:"dnsServers"`
}
