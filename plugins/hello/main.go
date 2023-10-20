package main

import (
	"fmt"
	"log"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const Version = "1.0.0"
const Priority = 2

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
	host, err := kong.Request.GetHeader("host")
	if err != nil {
		log.Printf("Error reading 'host' header: %s", err.Error())
	}

	message := conf.Message
	if message == "" {
		message = "hello"
	}
	kong.Response.SetHeader("x-hello-from-go", fmt.Sprintf("Go says %s to %s", message, host))
}
