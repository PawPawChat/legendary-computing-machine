package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	flag.Set("env", "testing")
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Env().ServerAddr)
	assert.NotEmpty(t, config.Env().LogLevel)
}

func TestLoadDefaultConfig(t *testing.T) {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := LoadDefaultConfig()
	assert.NoError(t, err)

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Env().ServerAddr)
	assert.NotEmpty(t, config.Env().LogLevel)
}

func TestConfigureLogger(t *testing.T) {
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config, config.Env().ServerAddr)
	assert.NotNil(t, config, config.Env().LogLevel)

	assert.NoError(t, ConfigureLogger(config))
}
