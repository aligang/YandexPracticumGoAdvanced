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
}

func (a *Agent) Do(r *http.Request) (*http.Response, error) {
	fun := a.client.Do
	if a.cryptoPlugin != nil {
		fun = a.cryptoPlugin.EncryptWithPublicKey(fun)
	}
	return fun(r)
}

func New(conf *config.AgentConfig) *Agent {
	a := &Agent{
		conf: conf,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	if conf.CryptoKey != "" {
		a.cryptoPlugin = encrypt.GetAgentPlugin(conf)
	}
	return a
}
