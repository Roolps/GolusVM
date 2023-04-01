package golusvm

import (
	"io"
	"net/http"
	"net/url"
)

// Makes post request to the api...
func (c *ApiClient) request(action string, fields map[string]string) *[]byte {
	payload := url.Values{
		"key":    []string{c.Key},
		"id":     []string{c.ID},
		"action": []string{action},
		"rdtype": []string{"json"},
	}
	for key, val := range fields {
		payload.Add(key, val)
	}

	client := &http.Client{}

	resp, err := client.PostForm(c.Endpoint, payload)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	// Return raw so func can marshal into correct type
	return &raw
}
