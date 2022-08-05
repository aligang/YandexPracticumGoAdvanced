package storage

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"strconv"
)

type Storage struct {
	dbGauge   map[string]float64
	dbCounter map[string]int64
}

func (s *Storage) init() {
	s.dbGauge = map[string]float64{}
	s.dbCounter = map[string]int64{}
}

func New() *Storage {
	s := &Storage{}
	s.init()
	return s
}

func (s *Storage) Update(metric *metric.Metric) {
	switch metricType := metric.MetricType; metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metric.MetricValue, 64)
		if err == nil {
			s.dbGauge[metric.MetricName] = value
			fmt.Printf("%s: %f\n", metric.MetricName, value)
		} else {
			fmt.Println(err)
		}
	case "counter":
		value, err := strconv.ParseInt(metric.MetricValue, 10, 64)
		if err == nil {
			s.dbCounter[metric.MetricName] += value
			fmt.Printf("%s: %d\n", metric.MetricName, s.dbCounter[metric.MetricName])
		} else {
			fmt.Println(err)
		}
	}

}
