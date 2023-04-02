package golusvm

import (
	"encoding/json"
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
	temp := point(strings.Split(raw[key], ","))
	return temp
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
