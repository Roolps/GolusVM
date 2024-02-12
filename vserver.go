package golusvm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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
