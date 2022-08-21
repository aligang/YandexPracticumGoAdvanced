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

func PushData(client *http.Client, m *metric.Metrics) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/update/", buf)
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

func PullData(client *http.Client, m *metric.Metrics) (metric.Metrics, error) {
	buf := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(buf)
	err := jsonEncoder.Encode(*m)
	if err != nil {
		fmt.Println("Error During serialization ")
		panic(err)
	}
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/value/", buf)
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
		fmt.Println("Value was provided")
		//fmt.Println(*pulledMetric.Delta)
		return pulledMetric, nil
	} else {
		fmt.Println("No value were provided")
		return pulledMetric, errors.New("")
	}
}

//func ComposeURI(typeName string, fieldName string, value string) string {
//	return fmt.Sprintf("http://127.0.0.1:8080/update/%s/%s/%s", typeName, fieldName, value)
//}

func SendMetrics(reportInterval int, stats *metric.Stats) {
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
			PushData(client, m)
		}
		for name, value := range stats.Counter {
			m := &metric.Metrics{ID: name, MType: "counter", Delta: &value}
			fmt.Printf("Checking old value of counter: %s\n", name)
			fetchedMetric, err := PullData(client, m)
			//fmt.Println(fetchedMetric)
			//fmt.Println(err)
			if err == nil {
				fmt.Printf("counter: %s=%d\n", name, *fetchedMetric.Delta)
			} else {
				fmt.Printf("Record for counter: %s was not found\n", name)
			}
			fmt.Printf("Updating value of counter: %s=%d\n", name, *m.Delta)
			PushData(client, m)
			fmt.Printf("Checking new value of counter: %s\n", name)
			fetchedMetric, err = PullData(client, m)
			if err == nil {
				fmt.Printf("counter: %s=%d", name, *fetchedMetric.Delta)
			} else {
				fmt.Printf("Record for counter: %s was not found\n", name)
			}
		}
		iteration++
	}
}
