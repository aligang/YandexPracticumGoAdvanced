package storage

import (
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/database"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/storage/memory"
)

func New(conf *config.ServerConfig) (Storage, string) {
	switch true {
	case len(conf.DatabaseDsn) > 0:
		logging.Debug("Configuring SQL Database Storage")
		return database.New(conf), "Database"
	case len(conf.StoreFile) > 0:
		logging.Debug("Configuring In-Memory Storage")
		return memory.New(conf), "Memory"
	default:
		panic("Unsupported storage configuration")
	}
	return nil, ""
}

type Storage interface {
	BulkUpdate(metrics map[string]metric.Metrics)
	Dump() metric.MetricMap
	Get(metricName string) (metric.Metrics, bool)
	Update(metrics metric.Metrics)
	IsAlive() error
}
