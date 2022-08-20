package storage

import (
	"fmt"
	"strconv"
	"sync"
)

type gauge map[string]float64
type counter map[string]int64

type Storage struct {
	DBGauge     gauge
	DBCounter   counter
	GaugeLock   sync.Mutex
	CounterLock sync.Mutex
}

func (s *Storage) init() {
	s.DBGauge = gauge{}
	s.DBCounter = counter{}
	s.GaugeLock = sync.Mutex{}
	s.CounterLock = sync.Mutex{}
}

func Define(gaugeDB gauge, counterDB counter) *Storage {
	s := &Storage{}
	s.DBCounter = counterDB
	s.DBGauge = gaugeDB
	s.GaugeLock = sync.Mutex{}
	s.CounterLock = sync.Mutex{}
	return s
}

func New() *Storage {
	s := &Storage{}
	s.init()
	return s
}

func (s *Storage) UpdateGauge(metricName string, metricValue float64) {
	s.DBGauge[metricName] = metricValue
}

func (s *Storage) UpdateCounter(metricName string, metricValue int64) {
	s.DBCounter[metricName] = metricValue
}

func (s *Storage) Get(metricType, metricName string) (any, bool) {
	var value any
	var found bool
	switch metricType {
	case "gauge":
		value, found = s.DBGauge[metricName]
	case "counter":
		value, found = s.DBCounter[metricName]
	}
	return value, found
}

func (s *Storage) Dump() string {
	result := ""
	for k, v := range s.DBGauge {
		result = result + fmt.Sprintf("%s     %s\n", k, strconv.FormatFloat(v, 'f', -1, 64))
	}
	for k, v := range s.DBCounter {
		result = result + fmt.Sprintf("%s     %d\n", k, v)
	}
	return result
}
