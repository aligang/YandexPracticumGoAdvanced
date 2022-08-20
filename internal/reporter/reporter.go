package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
	"time"
)

func MakeCall(client *http.Client, m *metric.Metrics) {
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

	if err != nil {
		fmt.Println("Error During Sending request ")
		panic(err)
	}

	defer response.Body.Close()
}

//func ComposeURI(typeName string, fieldName string, value string) string {
//	return fmt.Sprintf("http://127.0.0.1:8080/update/%s/%s/%s", typeName, fieldName, value)
//}

func SendMetrics(stats *metric.Stats) {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		//s := reflect.ValueOf(*stats)
		//for i := 0; i < s.NumField(); i++ {
		//	e := s.Field(i)
		//	for j := 0; j < e.NumField(); j++ {
		//		metricValue := e.Field(j)
		//		metricName := e.Type().Field(j).Name
		//		metricType := strings.ToLower(metricValue.Type().Name())
		//		value := fmt.Sprintf("%v", metricValue)
		//		//url := ComposeURI(metricType, metricName, value)
		//		MakeCall(client, metric.Metrics{ID: metricName, MType: metricType, Value: float64(1000)})
		//	}
		//}
		for name, value := range stats.Gauge {
			MakeCall(client, &metric.Metrics{ID: name, MType: "Gauge", Value: &value})
		}
		for name, value := range stats.Counter {
			MakeCall(client, &metric.Metrics{ID: name, MType: "Counter", Delta: &value})
		}
	}
}
