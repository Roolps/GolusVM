package golusvm

import (
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
