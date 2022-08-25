package reporter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"io"
	"net/http"
	"time"
)

func PushData(address string, client *http.Client, m *metric.Metrics) {
	jbuf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jbuf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		panic(err)
	}
	URI := fmt.Sprintf("http://%s/update/", address)
	gbuf := &bytes.Buffer{}
	gz, err := gzip.NewWriterLevel(gbuf, gzip.BestSpeed)
	if err != nil {
		fmt.Println("Error During compressor creation")
		panic(err)
	}
	res, _ := io.ReadAll(jbuf)
	if err != nil {
		fmt.Println("Error During fetching data for compressiong")
		panic(err)
	}
	gz.Write(res)
	gz.Close()

	if err != nil {
		fmt.Println("Error During compression")
		panic(err)
	}
	request, err := http.NewRequest("POST", URI, gbuf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		fmt.Println("Error During communication ")
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Encoding", "gzip")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error During Pushing data ")
	} else {
		defer response.Body.Close()
	}

}

func PullData(address string, client *http.Client, m *metric.Metrics) (metric.Metrics, error) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		panic(err)
	}
	URI := fmt.Sprintf("http://%s/value/", address)
	request, err := http.NewRequest("POST", URI, buf)
	fmt.Printf("Seding request to: URI: %s\n", URI)
	if err != nil {
		fmt.Println("Error During building ")
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	var pulledMetric metric.Metrics
	if err != nil {
		fmt.Println("Error During Pulling data ")
		return pulledMetric, errors.New("")
	} else {
		defer response.Body.Close()
		if response.StatusCode == 200 {
			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(&pulledMetric)
			if err != nil {
				fmt.Println("Error During Parsing data ")
				panic(err)
			}
			return pulledMetric, nil
		} else {
			return pulledMetric, errors.New("")
		}
	}
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
			PushData(agentConfig.Address, client, m)
		}
		for name, value := range stats.Counter {
			m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
			//fmt.Printf("Checking old value of counter: %s\n", name)
			//fetchedMetric, err := PullData(agentConfig.Address, client, m)
			//if err == nil {
			//	fmt.Printf("counter: %s=%d\n", name, *fetchedMetric.Delta)
			//} else {
			//	fmt.Printf("Record for counter: %s was not found\n", name)
			//}
			fmt.Printf("Updating value of counter: %+v\n", *m)
			PushData(agentConfig.Address, client, m)
			//fmt.Printf("Checking new value of counter: %s\n", name)
			//fetchedMetric, err = PullData(agentConfig.Address, client, m)
			//if err == nil {
			//	fmt.Printf("counter: %s=%d", name, *fetchedMetric.Delta)
			//} else {
			//	fmt.Printf("Record for counter: %s was not found\n", name)
			//}
		}
		iteration++
	}
}
