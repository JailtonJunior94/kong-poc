package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 1

func main() {
	server.StartServer(New, Version, Priority)
}

type Config struct {
	Message string `json:"message"`
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	apiKey, err := kong.Request.GetHeader("api-key")
	if err != nil {
		log.Printf("Error reading 'API KEY' header: %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, "https://6ectb62ojh.execute-api.us-east-1.amazonaws.com/api", nil)
	if err != nil {
		log.Printf("Error mount req: %s", err.Error())
	}
	req.Header.Set("api-key", apiKey)

	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error mount res: %s", err.Error())
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		kong.Response.SetStatus(http.StatusUnauthorized)
		return
	}

	kong.Response.SetStatus(http.StatusOK)
}
