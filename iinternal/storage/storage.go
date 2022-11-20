package storage

import (
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/storage/database"
	"github.com/aligang/YandexPracticumGoAdvanced/iinternal/storage/memory"
)

func New(conf *config.ServerConfig) (Storage, string) {
	var storage Storage
	var storageType string
	switch true {
	case len(conf.DatabaseDsn) > 0:
		logging.Debug("Configuring SQL Database Storage")
		storage = database.New(conf)
		storageType = "Database"
	default:
		logging.Debug("Configuring In-Memory Storage")
		storage = memory.New(conf)
		storageType = "Memory"
	}
	logging.Debug("Succeed")
	return storage, storageType
}

type Storage interface {
	BulkUpdate(metrics map[string]metric.Metrics) error
	Dump() metric.MetricMap
	Get(metricName string) (metric.Metrics, bool)
	Update(metrics metric.Metrics) error
	IsAlive() error
}