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

func (c *ApiClient) CreateVirtualServer(s *CreateVirtualServer) *CreatedVirtualServer {
	log.Println("Creating virtual server...")
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

/*
Notes:
	- need to add verifications... type can either be openvz, xen, xen hvm or kvm...
	- One of node or nodegroup must be present
*/
