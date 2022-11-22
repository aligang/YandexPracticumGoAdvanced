package storage

import (
	"github.com/aligang/YandexPracticumGoAdvanced/lib/config"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/metric"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage/database"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/storage/memory"
)

// New fabric that returns repository struct
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

// Storage interface, that represents repository methods
type Storage interface {

	//BulkUpdate - add/update several metrics to repository
	BulkUpdate(metrics map[string]metric.Metrics) error

	//Dump - retrieve all metrics from repository
	Dump() metric.MetricMap

	//Get - retrieve single metric from repository
	Get(metricName string) (metric.Metrics, bool)

	//Update - add/update single metric to repository
	Update(metrics metric.Metrics) error

	//IsAlive - check repository reachablity
	IsAlive() error
}
