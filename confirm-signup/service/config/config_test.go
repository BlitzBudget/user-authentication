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

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)
	userPoolId := "userPoolId"
	// Set environment variable
	os.Setenv("USERPOOL_ID", userPoolId)
	// returns value of USERPOOL_ID
	fmt.Println("USERPOOL_ID:", os.Getenv("USERPOOL_ID"))
	UserPoolId = os.Getenv("USERPOOL_ID")

	assert.Equal(userPoolId, UserPoolId)
}
