package storage

import (
	"fmt"
	"strconv"
)

type gauge map[string]float64
type counter map[string]int64

type Storage struct {
	dbGauge   gauge
	dbCounter counter
}

func (s *Storage) init() {
	s.dbGauge = gauge{}
	s.dbCounter = counter{}
}

func Define(gaugeDB gauge, counterDB counter) *Storage {
	s := &Storage{}
	s.dbCounter = counterDB
	s.dbGauge = gaugeDB
	return s
}

func New() *Storage {
	s := &Storage{}
	s.init()
	return s
}

func (s *Storage) Update(metricType, metricName, metricValue string) {
	switch metricType {
	case "gauge":
		value, err := strconv.ParseFloat(metricValue, 64)
		if err == nil {
			s.dbGauge[metricName] = value
		} else {
			fmt.Println(err)
		}
	case "counter":
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err == nil {
			s.dbCounter[metricName] += value
		} else {
			fmt.Println(err)
		}
	}
}

func (s *Storage) Get(metricType, metricName string) (string, bool) {
	var result string
	var found bool
	var gauge float64
	var counter int64
	switch metricType {
	case "gauge":
		gauge, found = s.dbGauge[metricName]
		if !found {
			result = ""
		} else {
			result = strconv.FormatFloat(gauge, 'f', -1, 64)
		}

	case "counter":
		counter, found = s.dbCounter[metricName]
		if !found {
			result = ""
		} else {
			result = fmt.Sprintf("%d", counter)
		}

	}
	return result, found
}

func (s *Storage) Dump() string {
	result := ""
	for k, v := range s.dbGauge {
		result = result + fmt.Sprintf("%s     %s\n", k, strconv.FormatFloat(v, 'f', -1, 64))
	}
	for k, v := range s.dbCounter {
		result = result + fmt.Sprintf("%s     %d\n", k, v)
	}
	return result
}
