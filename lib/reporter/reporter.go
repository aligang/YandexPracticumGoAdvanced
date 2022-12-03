package reporter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
)

func PushData(address string, client *http.Client, m *metric.Metrics) error {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		logging.Logger.Warn().Msg("Error During serialization ")
		fmt.Println()
		return err
	}
	URI := fmt.Sprintf("http://%s/update/", address)
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
	request, err := http.NewRequest("POST", URI, gbuf)
	logging.Debug("Sending request to: URI: %s\n", URI)
	if err != nil {
		logging.Logger.Warn().Msg("Error During communication ")
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := client.Do(request)

	if err != nil {
		logging.Warn("Error During Pushing data ")
		return err
	}
	defer response.Body.Close()
	return nil
}

func BulkPushData(address string, client *http.Client, m []metric.Metrics) error {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(m)
	if err != nil {
		logging.Warn("Error During serialization ")
		return err
	}
	URI := fmt.Sprintf("http://%s/update/", address)
	gbuf := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
	if err != nil {
		logging.Logger.Warn().Msg("Error During compressor creation")
		return err
	}
	res, err := io.ReadAll(jbuf)
	logging.Debug("Going to send json: %s", string(res))
	if err != nil {
		logging.Logger.Warn().Msg("Error During fetching data for compression")
		return err
	}
	_, err = gz.Write(res)
	gz.Close()

	if err != nil {
		logging.Warn("Error During compression")
		return err
	}
	request, err := http.NewRequest("POST", URI, gbuf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		logging.Logger.Warn().Msg("Error During communication ")
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := client.Do(request)

	if err != nil {
		logging.Logger.Warn().Msg("Error During Pushing data ")
		return err
	}
	defer response.Body.Close()
	return nil
}

// SendMetrics encode and send metrics one-by-one
func SendMetrics(agentConfig *config.AgentConfig, bus <-chan metric.Stats) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	iteration := 0
	for stats := range bus {
		logging.Debug("Running Iteration %d\n", iteration)
		for name, value := range stats.Gauge {
			m := &metric.Metrics{ID: name, MType: "gauge", Value: &value}
			if len(agentConfig.Key) > 0 {
				hash.AddHashInfo(m, agentConfig.Key)
			}
			err := PushData(agentConfig.Address, client, m)
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
			err := PushData(agentConfig.Address, client, m)
			if err != nil {
				logging.Logger.Warn().Msg(err.Error())
			}
		}
		iteration++
	}
}

// BulkSendMetrics bulk encode and send several metrics
func BulkSendMetrics(agentConfig *config.AgentConfig, bus <-chan metric.Stats) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	iteration := 0
	for stats := range bus {
		metrics := []metric.Metrics{}
		logging.Debug("Running Iteration %d\n", iteration)
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
		err := BulkPushData(agentConfig.Address, client, metrics)
		if err != nil {
			logging.Warn("Could not Push data: %s", err.Error())
		}
		iteration++
	}
}