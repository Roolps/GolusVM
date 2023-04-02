package golusvm

type ServerInformation struct {
	CTID_XID string `json:"ctid-xid"`

	// Ip address of returned server
	IP string `json:"ipaddress"`

	// Hostname of returned server
	Hostname string `json:"hostname"`

	// Template of returned server
	Template string `json:"template"`

	// HDD in Bytes of returned server
	HDD int `json:"hdd,string"`

	// Memory in Bytes of returned server
	Memory int `json:"memory,string"`

	// Swap burst in Bytes of returned server
	SwapBurst int `json:"swap-burst,string"`

	// Type of returned server
	Type string `json:"type"`

	// CPUs assigned to returned server
	CPUs string `json:"cpus"`

	// Mac address of returned server
	Mac string `json:"mac"`
}

// NIL GETS OMMITTED SO USE POINTER...
type CreateVirtualServer struct {
	// options: openvz, xen, xen hvm, kvm
	Type string `json:"type"`

	// Optional if nodegroup is set - Node ID
	Node int `json:"node,omitempty,string"`

	// Optional if node is set - Node group ID
	NodeGroup int `json:"nodegroup,string"`

	// Hostname of the VPS
	HostName string `json:"hostname"`

	// Root password (optional for custom password)
	Password string `json:"password,omitempty"`

	// Client Username, if not set defaults to system admin user
	Username string `json:"username"`

	// Plan name
	Plan string `json:"plan"`

	// Template or ISO name
	Template string `json:"template"`

	// Amount if IPv4 addresses
	IPS int `json:"ips,string"`

	// True or false
	RandomIPv4 bool `json:"randomipv4,omitempty,string"`

	// Can be 1 or 0 - This allows to to define templates & isos for Xen HVM
	HMVT int `json:"hmvt,omitempty,string"`

	// Set to overide plan memory
	CustomMemory int `json:"custommemory,omitempty,string"`

	// Set to override plan diskspace
	CustomDiskspace int `json:"customdiskspace,omitempty,string"`

	// Set to override plan bandwidth
	CustomBandwidth int `json:"custombandwidth,omitempty,string"`

	// Set to override plan cpu
	CustomCPU int `json:"customcpu,omitempty,string"`

	// Set to add amount of extra ips
	CustomExtraIP int `json:"customextraip,omitempty,string"`

	// Optional, 1 or 2, 1 = cpanel monthly, 2 = cpanel yearly
	IssueLicense int `json:"issuelicense,omitempty,string"`

	// Optional whether to assign an internal IP
	InternalIP bool `json:"internalip,omitempty,string"`
}

type CreatedVirtualServer struct {
	// Main server IP Address!
	MainIPAddress string `json:"mainipaddress"`

	// Array of extra IP addresses
	ExtraIPs *[]string

	// Root password of the created machine
	RootPassword string `json:"rootpassword"`

	// VserverID
	VServerID int `json:"vserverid,string"`

	// ConsoleUser for the new server
	ConsoleUser string `json:"consoleuser"`

	// Console password for the new server
	ConsolePassword string `json:"consolepassword"`

	// Hostname of the new server
	HostName string `json:"hostname"`

	// VM Id of the created server
	VirtID string `json:"virtid"`

	// Node ID of the created server
	NodeID string `json:"nodeid"`
}
