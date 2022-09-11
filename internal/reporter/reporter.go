package reporter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/hash"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"io"
	"net/http"
	"time"
)

func PushData(address string, client *http.Client, m *metric.Metrics) error {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		return err
	}
	URI := fmt.Sprintf("http://%s/update/", address)
	gbuf := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
	if err != nil {
		fmt.Println("Error During compressor creation")
		return err
	}
	res, err := io.ReadAll(jbuf)
	if err != nil {
		fmt.Println("Error During fetching data for compressiong")
		return err
	}
	_, err = gz.Write(res)
	gz.Close()

	if err != nil {
		fmt.Println("Error During compression")
		return err
	}
	request, err := http.NewRequest("POST", URI, gbuf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		fmt.Println("Error During communication ")
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error During Pushing data ")
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
		fmt.Println("Error During serialization ")
		return err
	}
	URI := fmt.Sprintf("http://%s/update/", address)
	gbuf := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
	if err != nil {
		fmt.Println("Error During compressor creation")
		return err
	}
	res, err := io.ReadAll(jbuf)
	if err != nil {
		fmt.Println("Error During fetching data for compressiong")
		return err
	}
	_, err = gz.Write(res)
	gz.Close()

	if err != nil {
		fmt.Println("Error During compression")
		return err
	}
	request, err := http.NewRequest("POST", URI, gbuf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		fmt.Println("Error During communication ")
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error During Pushing data ")
		return err
	}
	defer response.Body.Close()
	return nil
}

func PullData(address string, client *http.Client, m *metric.Metrics) (*metric.Metrics, error) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		return nil, err
	}
	URI := fmt.Sprintf("http://%s/value/", address)
	request, err := http.NewRequest("POST", URI, buf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		fmt.Println("Error During building ")
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	var pulledMetric metric.Metrics
	if err != nil {
		fmt.Println("Error During Pulling data ")
		return &pulledMetric, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("")
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&pulledMetric)
	if err != nil {
		fmt.Println("Error During Parsing data ")
		return nil, err
	}
	return &pulledMetric, nil
}

func SendMetrics(agentConfig *config.AgentConfig, stats *metric.Stats) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	ticker := time.NewTicker(agentConfig.ReportInterval)
	iteration := 0
	for {
		<-ticker.C
		fmt.Printf("Running Iteration %d\n", iteration)
		for name, value := range stats.Gauge {
			m := &metric.Metrics{ID: name, MType: "gauge", Value: &value}
			if len(agentConfig.Key) > 0 {
				hash.AddHashInfo(m, agentConfig.Key)
			}
			err := PushData(agentConfig.Address, client, m)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		for name, value := range stats.Counter {
			m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
			if len(agentConfig.Key) > 0 {
				hash.AddHashInfo(m, agentConfig.Key)
			}
			fmt.Printf("Updating value of counter: %+v with delta: %d\n", *m, *m.Delta)
			err := PushData(agentConfig.Address, client, m)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		iteration++
	}
}

func BulkSendMetrics(agentConfig *config.AgentConfig, stats *metric.Stats) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	ticker := time.NewTicker(agentConfig.ReportInterval)
	iteration := 0
	for {
		metrics := []metric.Metrics{}
		<-ticker.C
		fmt.Printf("Running Iteration %d\n", iteration)
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
			fmt.Printf("Updating value of counter: %+v with delta: %d\n", *m, *m.Delta)
			metrics = append(metrics, *m)
		}
		err := BulkPushData(agentConfig.Address, client, metrics)
		if err != nil {
			fmt.Println("Could not Push data")
			fmt.Println(err.Error())
		}
		iteration++
	}
}
