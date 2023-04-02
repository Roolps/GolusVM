package golusvm

// Node statistics
type NodeStatistics struct {
	// Node ID
	ID int `json:"id,string"`

	// Node virtualisation
	Virt string `json:"virt"`

	// Node name
	Name string `json:"name"`

	// Node main IP
	IP string `json:"ip"`

	// Node memory limit
	MemoryLimit int `json:"memorylimit"`

	// Node disk limit
	DiskLimit int `json:"disklimit"`

	// Node hostname
	HostName string `json:"hostname"`

	// Node country
	Country string `json:"country"`

	// Node city
	City string `json:"city"`

	// SSH Port
	SSHPort int `json:"sshport"`

	// Node Arch
	Arch string `json:"arch"`

	// Free memory
	FreeMemory int `json:"freememory,string"`

	// Allocated memory
	AllocatedMemory int `json:"allocatedmemory"`

	// Allocated Bandwidth
	AllocatedBandwidth int `json:"allocatedbandwidth"`

	// Free disk space
	FreeDisk int `json:"freedisk,string"`

	// Free IPv6 Addresses
	FreeIPv6 int `json:"freeipv6,string"`

	// Free IPv6 Subnets
	FreeIPv6Subnets int `json:"freeipv6subnets,string"`

	// Node group ID
	NodeGroupID int `json:"nodegroupid,string"`

	// Node group name
	NodeGroupName string `json:"nodegroupname"`

	// Number of virtual servers
	VirtualServers int `json:"virtualservers"`

	// Number of free ips
	FreeIPs int `json:"freeips"`
}

type TemplateList struct {
	Templates    *[]string
	TemplatesHVM *[]string
	TemplatesKVM *[]string
}

type Resources struct {
	FreeHDD    int `json:"freehdd,string"`
	FreeMemory int `json:"freememory,string"`
}
