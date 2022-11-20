package memory

import (
	"sync"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
)

type MemStorage struct {
	Metrics      metric.MetricMap
	Lock         sync.Mutex
	BackupConfig BackupConfig
	Type         string
}

func (s *MemStorage) Init(conf *config.ServerConfig) {
	s.Metrics = metric.MetricMap{}
	s.Lock = sync.Mutex{}
	s.Type = "memory"
	if conf != nil {
		if conf.Restore {
			s.Restore(conf)
		}
		s.ConfigureBackup(conf)
	}
}

func (s *MemStorage) Load(metrics metric.MetricMap) {
	s.Metrics = metrics
}

func New(conf *config.ServerConfig) *MemStorage {
	s := &MemStorage{}
	s.Init(conf)
	return s
}

func (s *MemStorage) Update(metrics metric.Metrics) error {

	if metrics.MType == "gauge" {
		s.Metrics[metrics.ID] = metrics
	}
	if metrics.MType == "counter" {
		if _, exists := s.Metrics[metrics.ID]; !exists {
			s.Metrics[metrics.ID] = metrics
		} else {
			value := *s.Metrics[metrics.ID].Delta + *metrics.Delta
			*s.Metrics[metrics.ID].Delta = value
		}
	}
	if s.BackupConfig.enable && !s.BackupConfig.Periodic {
		logging.Debug("Staring On-Demand Backup")
		BackupDo(s)
	}
	return nil
}

func (s *MemStorage) BulkUpdate(metrics map[string]metric.Metrics) error {
	for _, m := range metrics {
		s.Update(m)
	}
	return nil
}

func (s *MemStorage) Get(metricName string) (metric.Metrics, bool) {
	value, found := s.Metrics[metricName]
	return value, found
}

func (s *MemStorage) Dump() metric.MetricMap {
	return s.Metrics
}

func (s *MemStorage) IsAlive() error {
	return nil
}
