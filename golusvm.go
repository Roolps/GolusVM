package golusvm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type APIClient struct {
	Key      string
	ID       string
	Endpoint string
}

func (c *APIClient) request(method string, action string, fields map[string]string) ([]byte, error) {
	client := &http.Client{}

	payload := url.Values{
		"key":    []string{c.Key},
		"id":     []string{c.ID},
		"action": []string{action},
		"rdtype": []string{"json"},
	}
	for key, val := range fields {
		payload.Add(key, val)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%v?%v", c.Endpoint, payload.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := extractError(raw); err != nil {
		return nil, err
	}
	// log.Println(string(raw))
	return raw, nil
}

type apiError struct {
	Status string `xml:"status" json:"status"`
	Msg    string `xml:"statusmsg" json:"statusmsg"`
}

func extractError(raw []byte) error {
	e := &apiError{}
	if err := json.Unmarshal(raw, &e); err != nil {
		// not json - try XML
		rawstr := fmt.Sprintf("<error>%v</error>", string(raw))
		if err := xml.Unmarshal([]byte(rawstr), &e); err != nil {
			return fmt.Errorf("something went wrong while reading body: %v", err)
		}
	}
	if e.Status == "error" {
		return fmt.Errorf("error: %v", e.Msg)
	}
	return nil
}
