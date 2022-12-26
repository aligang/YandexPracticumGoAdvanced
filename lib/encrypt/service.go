package encrypt

import (
	"crypto/rsa"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
)

type EncryptionPlugin struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GetServerPlugin(c *config.ServerConfig) *EncryptionPlugin {
	return &EncryptionPlugin{
		PrivateKey: ReadPrivateKey(c.CryptoKey),
	}
}

func GetAgentPlugin(c *config.AgentConfig) *EncryptionPlugin {
	return &EncryptionPlugin{
		PublicKey: ReadPublicKey(c.CryptoKey),
	}
}
