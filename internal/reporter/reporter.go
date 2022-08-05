package reporter

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"net/http"
	"reflect"
	"time"
)

func SendCall(client *http.Client, typeName string, fieldName string, value string) {
	url := fmt.Sprintf("http://127.0.0.1:8080/update/%s/%s/%s", typeName, fieldName, value)
	fmt.Println(url)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error During building ")
		panic(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	_, err = client.Do(request)
	if err != nil {
		fmt.Println("Error During Sending request ")
		panic(err)
	}
}

func SendMetrics(stats *metric.Stats) {

	client := &http.Client{}

	for {
		time.Sleep(10 * time.Second)
		for metricName, metricType := range metric.Metrics() {
			v := reflect.ValueOf(*stats.MemStats).FieldByName(metricName)
			value := fmt.Sprintf("%v", v)
			SendCall(client, metricType, metricName, value)
		}

		v := reflect.ValueOf(*stats.OperStats)
		for i := 0; i < v.NumField(); i++ {
			metricType := v.Field(i).Type().Name()
			//typeName := v.Type().Field(i).Type
			metricName := v.Type().Field(i).Name
			value := fmt.Sprintf("%v", v.Field(i))
			//fmt.Println(value)
			SendCall(client, metricType, metricName, value)
		}
	}
}
