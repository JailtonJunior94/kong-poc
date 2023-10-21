package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const PluginName = "auth"
const Version = "0.1.0"
const Priority = 1000

const FailedResponse = `{"error": "%s is required"}`
const ReqFailedResponse = `{"error": "%s ReqFailedResponse"}`
const ErrMountReq = `{"error": "%s ErrMountReq"}`
const ErrStatusCodeReq = `{"error": "%s ErrStatusCodeReq"}`

type Config struct {
	HeaderKey string `json:"header_key"`
}

func main() {
	err := server.StartServer(New, Version, Priority)
	if err != nil {
		log.Fatalf("Failed start %s plugin", PluginName)
	}

}

func New() interface{} {
	return &Config{}
}

func (conf *Config) Access(kong *pdk.PDK) {
	headerKey, err := kong.Request.GetHeader(conf.HeaderKey)
	if err != nil {
		log.Printf("Error reading 'host' header: %s", err.Error())
	}

	headerResponse := make(map[string][]string, 0)
	headerResponse["Content-Type"] = []string{"application/json"}

	if headerKey == "" {
		kong.Response.Exit(400, []byte(fmt.Sprintf(FailedResponse, conf.HeaderKey)), headerResponse)
	}

	client := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "https://6ectb62ojh.execute-api.us-east-1.amazonaws.com/api", nil)
	if err != nil {
		log.Printf("Error mount req: %s", err.Error())
		kong.Response.Exit(http.StatusInternalServerError, []byte(fmt.Sprintf(ErrMountReq, conf.HeaderKey)), headerResponse)
	}
	req.Header.Set("api-key", headerKey)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error mount res: %s", err.Error())
		kong.Response.Exit(http.StatusInternalServerError, []byte(fmt.Sprintf(ReqFailedResponse, conf.HeaderKey)), headerResponse)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("Error mount res: %v", res.StatusCode)
		kong.Response.Exit(res.StatusCode, []byte(fmt.Sprintf(ErrStatusCodeReq, conf.HeaderKey)), headerResponse)
	}
}
