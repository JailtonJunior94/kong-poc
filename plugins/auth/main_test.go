package main

import (
	"net/http"
	"testing"

	"github.com/Kong/go-pdk/test"
	"github.com/stretchr/testify/assert"
)

func TestPluginConfig(t *testing.T) {
	env, err := test.New(t, test.Request{
		Method:  "GET",
		Url:     "http://example.com?q=search&x=9",
		Headers: map[string][]string{"header_key": {"45aabf67-8337-4eb7-8d2c-2cf6b554fbf4"}},
	})
	assert.NoError(t, err)

	env.DoHttps(&Config{HeaderKey: "header_key"})
	assert.Equal(t, http.StatusOK, env.ClientRes.Status)
}

func TestPluginWithInvalidHeaderKey(t *testing.T) {
	env, err := test.New(t, test.Request{
		Method:  "GET",
		Url:     "http://example.com?q=search&x=9",
		Headers: map[string][]string{"header_key": {"45aabf67-8337-4eb7-8d2c-2cf6b554fbf8"}},
	})
	assert.NoError(t, err)

	env.DoHttps(&Config{HeaderKey: "header_key"})
	assert.Equal(t, http.StatusUnauthorized, env.ClientRes.Status)
}
