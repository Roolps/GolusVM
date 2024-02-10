package golusvm

import (
	"net/http"
	"strconv"
)

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
