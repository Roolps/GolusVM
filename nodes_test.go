package golusvm

import (
	"log"
	"testing"
)

func TestListServers(t *testing.T) {
	newTest()
	nodes, err := testClient.ListNodesByID(KVM)
	if err != nil {
		t.Error(err)
	}
	log.Println(nodes)
	for _, n := range nodes {
		srvs, err := testClient.ListVirtualServers(n)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(srvs)
		}
	}
}
