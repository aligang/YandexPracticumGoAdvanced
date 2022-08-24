package storage

import (
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"sync"
)

type metricMap map[string]metric.Metrics

type Storage struct {
	Metrics      metricMap
	Lock         sync.Mutex
	BackupConfig BackupConfig
}

func (s *Storage) init() {
	s.Metrics = metricMap{}
	s.Lock = sync.Mutex{}
}

func (s *Storage) Load(metrics metricMap) {
	s.Metrics = metrics
}

func New() *Storage {
	s := &Storage{}
	s.init()
	return s
}

func (s *Storage) Update(metrics metric.Metrics) {

	if metrics.MType == "gauge" {
		s.Metrics[metrics.ID] = metrics
	}
	if metrics.MType == "counter" {
		if _, exists := s.Metrics[metrics.ID]; exists == false {
			s.Metrics[metrics.ID] = metrics
		} else {
			*s.Metrics[metrics.ID].Delta = *s.Metrics[metrics.ID].Delta + *metrics.Delta
		}
	}
	if s.BackupConfig.enable == true && s.BackupConfig.Periodic == false {
		fmt.Println("Staring On-Deman Backup")
		BackupDo(s)
	}
}

func (s *Storage) Get(metricName string) (metric.Metrics, bool) {
	var found bool
	value, found := s.Metrics[metricName]
	return value, found
}

func (s *Storage) Dump() metricMap {
	return s.Metrics
}
