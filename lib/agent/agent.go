package agent

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"net/http"
	"time"
)

type Agent struct {
	conf         *config.AgentConfig
	client       *http.Client
	cryptoPlugin *encrypt.EncryptionPlugin
	Do           func(r *http.Request) (*http.Response, error)
}

func New(conf *config.AgentConfig) *Agent {
	a := &Agent{
		conf: conf,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	a.Do = a.client.Do
	//compress.AgentCompression(a.Do)
	if conf.CryptoKey != "" {
		encrypt.GetAgentPlugin(conf).EncryptWithPublicKey(a.Do)
	}
	return a
}
