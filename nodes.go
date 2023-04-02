package golusvm

import (
	"encoding/json"
	"strconv"
)

func (c *ApiClient) ListISO(isoType string) *[]string {
	var iso *[]string
	raw := c.request("listiso", map[string]string{"type": "kvm"})
	iso = split(raw, "iso")
	return iso
}

type Resources struct {
	FreeHDD    int `json:"freehdd,string"`
	FreeMemory int `json:"freememory,string"`
}

func (c *ApiClient) XENNodeResources(nodeID int) Resources {
	var resources Resources
	raw := c.request("node-xenresources", map[string]string{"nodeid": strconv.Itoa(nodeID)})
	json.Unmarshal(*raw, &resources)
	return resources
}

func (c *ApiClient) ListNodeGroups(nodeType string) *[]string {
	var nodeGroups *[]string
	raw := c.request("listnodegroups", map[string]string{"type": nodeType})
	nodeGroups = split(raw, "nodegroups")
	return nodeGroups
}

func (c *ApiClient) ListNodes(nodeType string) *[]string {
	var nodeList *[]string
	raw := c.request("listnodes", map[string]string{"type": nodeType})
	nodeList = split(raw, "nodes")
	return nodeList
}

func (c *ApiClient) ListNodeIPs(nodeID int) *[]string {
	var nodeIpList *[]string
	raw := c.request("node-iplist", map[string]string{"nodeID": strconv.Itoa(nodeID)})
	nodeIpList = split(raw, "ips")
	return nodeIpList
}
