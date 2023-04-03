package golusvm

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type ApiClient struct {
	Key      string
	ID       string
	Endpoint string
}

func point[T any](t T) *T {
	return &t
}

func split(jsonRaw *[]byte, key string) *[]string {
	var raw map[string]string
	json.Unmarshal(*jsonRaw, &raw)
	temp := strings.Split(raw[key], ",")
	new := []string{}
	for _, val := range temp {
		new = append(new, strings.Replace(val, " ", "", -1))
	}
	return &new
}

func jsonMap(i interface{}) *map[string]string {
	b, _ := json.MarshalIndent(i, " ", "")
	x := &map[string]string{}
	json.Unmarshal(b, x)
	return x
}

func debug(b []byte) {
	x := &map[string]string{}
	json.Unmarshal(b, x)
	log.Println(x)
}

type apiStatus struct {
	Status    string `json:"status"`
	StatusMsg string `json:"statusmsg"`
}

// Extracts status message from api response
func extractStatus(raw *[]byte) error {
	status := apiStatus{}
	json.Unmarshal(*raw, &status)
	if status.Status == "error" {
		return errors.New(status.StatusMsg)
	}
	return nil
}
