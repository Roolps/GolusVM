package golusvm

import (
	"net/http"
	"testing"

	"github.com/joho/godotenv"
)

var testClient = &APIClient{}

func newTest() {
	envFile, _ := godotenv.Read(".env")
	testClient.ID = envFile["SOLUS_ID"]
	testClient.Key = envFile["SOLUS_KEY"]
	testClient.Endpoint = envFile["SOLUS_ENDPOINT"]
}

func TestRequestHandler(t *testing.T) {
	newTest()
	testClient.request(http.MethodGet, "vserver-info", map[string]string{"vserverid": "58"})
}
