package golusvm

import (
	"encoding/json"
	"log"
	"strconv"
)

func (c *ApiClient) GetServerInformation(sID int) *ServerInformation {
	var server *ServerInformation
	raw := c.request("vserver-info", map[string]string{"vserverid": strconv.Itoa(sID)})
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(*raw, &server)
	return server
}

func (c *ApiClient) GetServerState(sID int, nostatus bool, nograph bool) *VirtualServerState {
	var state *VirtualServerState
	raw := c.request("vserver-infoall", map[string]string{"vserverid": strconv.Itoa(sID), "nostatus": strconv.FormatBool(nostatus), "nograph": strconv.FormatBool(nograph)})
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(*raw, &state)
	state.IPAddresses = split(raw, "ipaddresses")
	state.IPv6Subnets = split(raw, "ipv6subnets")
	state.Bandwidth = split(raw, "bandwidth")
	state.HDD = split(raw, "hdd")
	state.Memory = split(raw, "memory")
	return state
}

func (c *ApiClient) CreateVirtualServer(s *CreateVirtualServer) *CreatedVirtualServer {
	var server *CreatedVirtualServer
	raw := c.request("vserver-create", *jsonMap(*s))
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(*raw, &server)
	return server
}

func (c *ApiClient) ChangeHostname(sID int, hostname string) string {
	raw := c.request("vserver-hostname", map[string]string{"vserverid": strconv.Itoa(sID), "hostname": hostname})
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return ""
	}
	return hostname
}

func (c *ApiClient) ReconfigureNetwork(sID int) {
	raw := c.request("vserver-reconfigure-network", map[string]string{"vserverid": strconv.Itoa(sID)})
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
	}
}

func (c *ApiClient) RebuildVirtualServer(sID int, template string) {
	raw := c.request("vserver-rebuild", map[string]string{"vserverid": strconv.Itoa(sID), "template": template})
	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
	}
}

func (c *ApiClient) ListVirtualServers(nID int) *[]ServerInformation {
	type temp struct {
		VirtualServers []ServerInformation
	}
	var b temp
	raw := c.request("node-virtualservers", map[string]string{"nodeid": strconv.Itoa(nID)})
	json.Unmarshal(*raw, &b)
	return &b.VirtualServers
}

/*
Notes:
	- need to add verifications... type can either be openvz, xen, xen hvm or kvm...
	- One of node or nodegroup must be present
*/
