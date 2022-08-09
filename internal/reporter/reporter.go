package reporter

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func MakeCall(client *http.Client, uri string) {
	request, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		fmt.Println("Error During building ")
		panic(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	response, err := client.Do(request)

	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error During Sending request ")
		panic(err)
	}

}

func ComposeURI(typeName string, fieldName string, value string) string {
	return fmt.Sprintf("http://127.0.0.1:8080/update/%s/%s/%s", typeName, fieldName, value)
}

func SendMetrics(stats *metric.Stats) {

	client := &http.Client{}

	for {
		time.Sleep(2 * time.Second)

		s := reflect.ValueOf(*stats)
		for i := 0; i < s.NumField(); i++ {
			e := s.Field(i)
			for j := 0; j < e.NumField(); j++ {
				metricValue := e.Field(j)
				metricName := e.Type().Field(j).Name
				metricType := strings.ToLower(metricValue.Type().Name())
				value := fmt.Sprintf("%v", metricValue)
				url := ComposeURI(metricType, metricName, value)
				MakeCall(client, url)
			}
		}
	}
}
