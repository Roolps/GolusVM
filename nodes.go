package golusvm

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (c *APIClient) ListNodesByID(virt Virtualization) ([]int, error) {
	raw, err := c.request(http.MethodGet, "node-idlist", map[string]string{"type": string(virt)})
	if err != nil {
		return nil, err
	}
	data := struct {
		Nodes string `json:"nodes"`
	}{}
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	nodesstr := strings.Split(data.Nodes, ",")
	nodes := []int{}
	for _, n := range nodesstr {
		i, _ := strconv.Atoi(n)
		nodes = append(nodes, i)
	}
	return nodes, nil
}

func (c *APIClient) ListVirtualServers(nodeid int) ([]*VirtualServer, error) {
	raw, err := c.request(http.MethodGet, "node-virtualservers", map[string]string{"nodeid": strconv.Itoa(nodeid)})
	if err != nil {
		return nil, err
	}
	data := struct {
		VirtualServers []*VirtualServer `json:"virtualservers"`
	}{
		VirtualServers: []*VirtualServer{},
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return data.VirtualServers, nil
}
