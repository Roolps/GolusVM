package golusvm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Virtualization string

const (
	OpenVZ Virtualization = "openvz"
	Xen    Virtualization = "xen"
	XenHVM Virtualization = "xen hvm"
	KVM    Virtualization = "kvm"
)

type VirtualServer struct {
	VServerID int    `json:"vserverid,string"`
	CTIDXID   string `json:"ctid-xid"`
	ClientID  int    `json:"clientid,string"`
	IPAddress string `json:"ipaddress"`
	Hostname  string `json:"hostname"`
	Template  string `json:"template"`

	HDD       int64 `json:"hdd,string"`
	Memory    int64 `json:"memory,string"`
	SwapBurst int64 `json:"swap-burst,string"`

	Type Virtualization `json:"type"`
	Mac  string         `json:"mac"`
}

type VirtualServerState struct {
	State         string         `json:"state"`
	IPs           []string       `json:"-"`
	IPsRAW        string         `json:"ipaddresses"`
	MainIPAddress string         `json:"mainipaddress"`
	Type          Virtualization `json:"type"`
	Node          string         `json:"node"`
}

type CreateVirtualServer struct {
	Type      Virtualization `json:"type"`
	NodeGroup int            `json:"nodegroup,string"`
	Hostname  string         `json:"hostname"`
	Password  string         `json:"password,omitempty"`
	Username  string         `json:"username"`
	Plan      string         `json:"plan"`
	Template  string         `json:"template"`

	IPs        int  `json:"ips,string"`
	RandomIPv4 bool `json:"randomipv4,omitempty"`

	HVM             int   `json:"hvmt,omitempty,string"`
	CustomMemory    int64 `json:"custommemory,omitempty,string"`
	CustomDiskSpace int64 `json:"customdiskspace,omitempty,string"`
	CustomBandwidth int64 `json:"custombandwidth,omitempty,string"`
	CustomCPU       int   `json:"customcpu,omitempty,string"`
	CustomExtraIP   int   `json:"customextraip,omitempty,string"`

	IssueLicense int `json:"issuelicense,omitempty,string"`
	InternalIP   int `json:"internalip,omitempty,string"`
}

type NewVirtualServer struct {
	VServerID int    `json:"vserverid,string"`
	VirtID    string `json:"virtid"`
	NodeID    int    `json:"nodeid"`

	MainIPAddress string   `json:"mainipaddress"`
	IPs           []string `json:"-"`
	IPsRAW        string   `json:"extraipaddress"`
	RootPassword  string   `json:"rootpassword"`
	Hostname      string   `json:"hostname"`
}

func (c *APIClient) CreateVirtualServer(s *CreateVirtualServer) (*NewVirtualServer, error) {
	raw, _ := json.Marshal(s)
	fields := map[string]string{}
	json.Unmarshal(raw, &fields)
	raw, err := c.request(http.MethodPost, "vserver-create", fields)
	if err != nil {
		return nil, err
	}
	srv := &NewVirtualServer{}
	if err := json.Unmarshal(raw, &srv); err != nil {
		return nil, err
	}
	return srv, nil
}

func (c *APIClient) VirtualServerState(id int) (*VirtualServerState, error) {
	raw, err := c.request(http.MethodGet, "vserver-infoall", map[string]string{"vserverid": strconv.Itoa(id), "nographs": "true"})
	if err != nil {
		return nil, err
	}
	server := &VirtualServerState{}
	if err := json.Unmarshal(raw, &server); err != nil {
		return nil, err
	}
	server.IPs = strings.Split(server.IPsRAW, ", ")
	return server, nil
}

// boot virtual server
func (c *APIClient) Boot(id int) error {
	_, err := c.request(http.MethodPost, "vserver-boot", map[string]string{"vserverid": strconv.Itoa(id)})
	return err
}

// shutdown virtual server
func (c *APIClient) Shutdown(id int) error {
	_, err := c.request(http.MethodPost, "vserver-shutdown", map[string]string{"vserverid": strconv.Itoa(id)})
	return err
}

// reboot virtual server
func (c *APIClient) Reboot(id int) error {
	_, err := c.request(http.MethodPost, "vserver-reboot", map[string]string{"vserverid": strconv.Itoa(id)})
	return err
}

// change virtual server hostname
func (c *APIClient) ChangeHostname(id int, hostname string) error {
	_, err := c.request(http.MethodPost, "vserver-hostname",
		map[string]string{
			"vserverid": strconv.Itoa(id),
			"hostname":  hostname,
		})
	if err != nil {
		return err
	}
	return nil
}

type state struct {
	State string `json:"state"`
}

func (c *APIClient) IsOnline(id int) (bool, error) {
	raw, err := c.request(http.MethodGet, "vserver-infoall", map[string]string{"vserverid": strconv.Itoa(id), "nographs": "true"})
	if err != nil {
		return false, err
	}
	s := &state{}
	json.Unmarshal(raw, &s)
	switch s.State {
	case "online":
		return true, nil
	case "offline":
		return false, nil
	}
	return false, fmt.Errorf("unknown state '%v'", s.State)
}

type rootPassword struct {
	Password string `json:"rootpassword"`
}

func (c *APIClient) ChangeRootPassword(id int) (string, error) {
	raw, err := c.request(http.MethodPost, "vserver-rootpassword", map[string]string{"vserverid": strconv.Itoa(id)})
	if err != nil {
		return "", err
	}
	p := &rootPassword{}
	json.Unmarshal(raw, &p)
	return p.Password, nil
}

func (c *APIClient) RebuildServer(id int, template string) error {
	_, err := c.request(http.MethodPost, "vserver-rebuild", map[string]string{"vserverid": strconv.Itoa(id), "template": template})
	return err
}
