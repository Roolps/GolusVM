package golusvm

import (
	"encoding/json"
	"log"
	"strconv"
)

/*---------------
  ISO FUNCTIONS
---------------*/

func (c *ApiClient) ListISO(isoType string) *[]string {
	var iso *[]string
	raw := c.request("listiso", map[string]string{"type": "kvm"})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	iso = split(raw, "iso")
	return iso
}

func (c *ApiClient) XENNodeResources(nodeID int) *Resources {
	var resources Resources
	raw := c.request("node-xenresources", map[string]string{"nodeid": strconv.Itoa(nodeID)})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	json.Unmarshal(*raw, &resources)
	return &resources
}

/*---------------
  GETTING NODES
---------------*/

// List nodes by name
func (c *ApiClient) ListNodesByName(nodeType string) *[]string {
	var nodeList *[]string
	raw := c.request("listnodes", map[string]string{"type": nodeType})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	nodeList = split(raw, "nodes")
	return nodeList
}

// List nodes by ID
func (c *ApiClient) ListNodesByID(nodeType string) *[]int {
	var nodeList []int
	raw := c.request("node-idlist", map[string]string{"type": nodeType})
	temp := split(raw, "nodes")
	for _, x := range *temp {
		b, _ := strconv.Atoi(x)
		nodeList = append(nodeList, b)
	}
	return &nodeList
}

// get statistics of a node
func (c *ApiClient) GetNodeStatistics(nodeID int) *NodeStatistics {
	var nodeStatistics NodeStatistics
	raw := c.request("node-statistics", map[string]string{"nodeid": strconv.Itoa(nodeID)})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	json.Unmarshal(*raw, &nodeStatistics)
	return &nodeStatistics
}

// List IPs assigned to a node
func (c *ApiClient) ListNodeIPs(nodeID int) *[]string {
	var nodeIpList *[]string
	raw := c.request("node-iplist", map[string]string{"nodeid": strconv.Itoa(nodeID)})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	nodeIpList = split(raw, "ips")
	return nodeIpList
}

// List available plans
func (c *ApiClient) ListPlans(planType string) *[]string {
	var plans *[]string
	raw := c.request("listplans", map[string]string{"type": planType})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	plans = split(raw, "plans")
	return plans
}

func (c *ApiClient) ListTemplates(templateType string, friendly bool) *TemplateList {
	var templateList TemplateList
	raw := c.request("listtemplates", map[string]string{"type": templateType, "listpipefriendly": strconv.FormatBool(friendly)})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	templateList.Templates = split(raw, "templates")
	templateList.TemplatesHVM = split(raw, "templateshvm")
	templateList.TemplatesKVM = split(raw, "templateskvm")

	return &templateList
}

func (c *ApiClient) ListNodeGroups(nodeType string) *[]string {
	var nodeGroups *[]string
	raw := c.request("listnodegroups", map[string]string{"type": nodeType})

	err := extractStatus(raw)
	if err != nil {
		log.Println(err)
		return nil
	}

	nodeGroups = split(raw, "nodegroups")
	return nodeGroups
}
