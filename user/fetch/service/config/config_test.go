package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)
	clientId := "clientId"
	// Set environment variable
	os.Setenv("CLIENT_ID", clientId)
	// returns value of CLIENT_ID
	fmt.Println("CLIENT_ID:", os.Getenv("CLIENT_ID"))
	ClientId = os.Getenv("CLIENT_ID")

	assert.Equal(clientId, ClientId)
}

func TestConfigClientSecret(t *testing.T) {
	assert := assert.New(t)
	clientSecret := "clientSecret"
	// Set environment variable
	os.Setenv("CLIENT_SECRET", clientSecret)
	// returns value of CLIENT_SECRET
	fmt.Println("CLIENT_SECRET:", os.Getenv("CLIENT_SECRET"))
	ClientSecret = os.Getenv("CLIENT_SECRET")

	assert.Equal(clientSecret, ClientSecret)
}
