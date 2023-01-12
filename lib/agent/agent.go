package agent

import (
	grpcService "github.com/aligang/YandexPracticumGoAdvanced/lib/app/grpc/generated/service"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type Agent struct {
	conf         *config.AgentConfig
	client       *http.Client
	grpcClient   grpcService.MetricsServiceClient
	cryptoPlugin *encrypt.EncryptionPlugin
	Do           func(r *http.Request) (*http.Response, error)
}

func New(conf *config.AgentConfig, grpcClient *grpc.ClientConn) *Agent {
	a := &Agent{
		conf: conf,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		grpcClient: grpcService.NewMetricsServiceClient(grpcClient),
	}
	a.Do = a.client.Do
	//if conf.CryptoKey != "" {
	//	a.Do = encrypt.GetAgentPlugin(conf).EncryptWithPublicKey(a.Do)
	//}
	return a
}
