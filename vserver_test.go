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

func TestBoot(t *testing.T) {
	newTest()
	err := testClient.Boot(58)
	if err != nil {
		t.Error(err)
	}
}

func TestShutdown(t *testing.T) {
	newTest()
	err := testClient.Shutdown(58)
	if err != nil {
		t.Error(err)
	}
}

func TestReboot(t *testing.T) {
	newTest()
	err := testClient.Reboot(58)
	if err != nil {
		t.Error(err)
	}
}
