package reporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
	"time"
)

func PushData(address string, client *http.Client, m *metric.Metrics) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		panic(err)
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s/update/", address), buf)
	if err != nil {
		fmt.Println("Error During building ")
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error During Pushing data ")
		panic(err)
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
	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s/value/", address), buf)
	if err != nil {
		fmt.Println("Error During building ")
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error During Pulling data ")
		panic(err)
	}
	var pulledMetric metric.Metrics
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

func SendMetrics(address string, reportInterval int, stats *metric.Stats) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	ticker := time.NewTicker(time.Second * time.Duration(reportInterval))
	iteration := 0
	for {
		<-ticker.C
		fmt.Printf("Running Iteration %d\n", iteration)
		for name, value := range stats.Gauge {
			m := &metric.Metrics{ID: name, MType: "gauge", Value: &value}
			PushData(address, client, m)
		}
		for name, value := range stats.Counter {
			m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
			fmt.Printf("Checking old value of counter: %s\n", name)
			fetchedMetric, err := PullData(address, client, m)
			if err == nil {
				fmt.Printf("counter: %s=%d\n", name, *fetchedMetric.Delta)
			} else {
				fmt.Printf("Record for counter: %s was not found\n", name)
			}
			fmt.Printf("Updating value of counter: %s=%d\n", name, *m.Delta)
			PushData(address, client, m)
			fmt.Printf("Checking new value of counter: %s\n", name)
			fetchedMetric, err = PullData(address, client, m)
			if err == nil {
				fmt.Printf("counter: %s=%d", name, *fetchedMetric.Delta)
			} else {
				fmt.Printf("Record for counter: %s was not found\n", name)
			}
		}
		iteration++
	}
}
