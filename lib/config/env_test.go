package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomAgentEnv(t *testing.T) {
	t.Run("AGENT ENV PARAMS", func(t *testing.T) {
		ref := &AgentConfig{Address: "127.0.0.0:12345", PollInterval: 20, ReportInterval: 100}
		test := &AgentConfig{}
		os.Setenv("ADDRESS", ref.Address)
		os.Setenv("POLL_INTERVAL", ref.PollInterval.String())
		os.Setenv("REPORT_INTERVAL", ref.ReportInterval.String())
		GetAgentENVConfig(test)
		assert.Equal(t, ref, test)
	})
}

func TestCustomServerEnv(t *testing.T) {
	t.Run("SERVER ENV PARAMS", func(t *testing.T) {
		ref := &ServerConfig{Address: "127.0.0.0:12345", StoreInterval: 1 * time.Second,
			Restore: true, StoreFile: "/abc"}
		test := &ServerConfig{}
		os.Setenv("ADDRESS", ref.Address)
		os.Setenv("STORE_INTERVAL", ref.StoreInterval.String())
		os.Setenv("RESTORE", "true")
		os.Setenv("STORE_FILE", ref.StoreFile)
		GetServerENVConfig(test)
		assert.Equal(t, ref, test)
	})
}
