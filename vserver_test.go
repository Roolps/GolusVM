package golusvm

import (
	"log"
	"testing"
)

func TestChangeHostnameWithInvalidHostname(t *testing.T) {
	newTest()
	err := testClient.ChangeHostname(58, "^^^")
	if err == nil {
		t.Error("expected error, got nil")
	} else {
		log.Println(err)
	}
}

func TestChangeHostnameWithValidHostname(t *testing.T) {
	newTest()
	err := testClient.ChangeHostname(58, "php.roolps.dev")
	if err != nil {
		t.Error(err)
	}
}
