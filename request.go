package golusvm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (c *ApiClient) Request(action string, fields map[string]string) map[string]string {
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
	var msg map[string]string
	raw, err := ioutil.ReadAll(resp.Body)
	log.Println(string(raw))
	json.Unmarshal(raw, &msg)

	if err != nil {
		return nil
	}

	// Marshal into JSON and return
	return msg
}
