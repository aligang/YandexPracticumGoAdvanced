package agent

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/encrypt"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"io"
	"net/http"
)

func (a *Agent) PushData(jbuf *bytes.Buffer, path string) error {
	URI := fmt.Sprintf("http://%s/%s/", a.conf.Address, path)
	gbuf := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
	if err != nil {
		logging.Logger.Warn().Msg("Error During compressor creation")
		return err
	}
	res, err := io.ReadAll(jbuf)
	if err != nil {
		logging.Logger.Warn().Msg("Error During fetching data for compressiong")
		return err
	}
	_, err = gz.Write(res)
	gz.Close()

	if err != nil {
		logging.Warn("Error During compression")
		return err
	}
	if a.conf.CryptoKey != "" {
		gbuf = encrypt.EncryptWithPublicKey(gbuf, a.cryptoPlugin.PublicKey)
	}

	request, err := http.NewRequest("POST", URI, gbuf)
	if err != nil {
		logging.Logger.Warn().Msg("Error During Request creation ")
		return err
	}
	logging.Debug("Sending request to: URI: %s\n", URI)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := a.Do(request)

	if err != nil {
		logging.Warn("Error During Pushing data ")
		return err
	}
	defer response.Body.Close()
	return nil
}

func (a *Agent) UnaryPushData(m *metric.Metrics) error {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		logging.Logger.Warn().Msg("Error During serialization ")
		fmt.Println()
		return err
	}
	err = a.PushData(jbuf, "update")
	if err != nil {
		return err
	}
	return nil
}

func (a *Agent) BulkPushData(m []metric.Metrics) error {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(m)
	if err != nil {
		logging.Warn("Error During serialization ")
		return err
	}
	err = a.PushData(jbuf, "updates")
	if err != nil {
		return err
	}
	return nil
}

// SendMetrics encode and send metrics one-by-one
func (a *Agent) SendMetrics(agentConfig *config.AgentConfig, bus <-chan metric.Stats, exit <-chan any) {
loop:
	for {
		select {
		case stats := <-bus:
			for name, value := range stats.Gauge {
				m := &metric.Metrics{ID: name, MType: "gauge", Value: &value}
				if len(agentConfig.Key) > 0 {
					hash.AddHashInfo(m, agentConfig.Key)
				}
				err := a.UnaryPushData(m)
				if err != nil {
					logging.Logger.Warn().Msg(err.Error())
				}
			}
			for name, value := range stats.Counter {
				m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
				if len(agentConfig.Key) > 0 {
					hash.AddHashInfo(m, agentConfig.Key)
				}
				logging.Debug("Updating value of counter: %+v with delta: %d\n", *m, *m.Delta)
				err := a.UnaryPushData(m)
				if err != nil {
					logging.Logger.Warn().Msg(err.Error())
				}
			}
		case <-exit:
			break loop
		}
	}
}

// BulkSendMetrics bulk encode and send several metrics
func (a *Agent) BulkSendMetrics(agentConfig *config.AgentConfig, bus <-chan metric.Stats, exit <-chan any) {
loop:
	for {
		select {
		case stats := <-bus:
			metrics := []metric.Metrics{}
			for name, value := range stats.Gauge {
				m := &metric.Metrics{ID: name, MType: "gauge", Value: &value}
				if len(agentConfig.Key) > 0 {
					hash.AddHashInfo(m, agentConfig.Key)
				}
				metrics = append(metrics, *m)
			}
			for name, value := range stats.Counter {
				m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
				if len(agentConfig.Key) > 0 {
					hash.AddHashInfo(m, agentConfig.Key)
				}
				logging.Debug("Updating value of counter: %+v with delta: %d\n", *m, *m.Delta)
				metrics = append(metrics, *m)
			}
			err := a.BulkPushData(metrics)
			if err != nil {
				logging.Warn("Could not Push data: %s", err.Error())
			}
		case <-exit:
			break loop
		}
	}
}
