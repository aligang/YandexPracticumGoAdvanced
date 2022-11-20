package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DBStorage struct {
	DB   *sql.DB
	Type string
}

func New(conf *config.ServerConfig) *DBStorage {
	db, err := sql.Open("pgx", conf.DatabaseDsn)
	if err != nil {
		panic(err)
	}
	s := &DBStorage{
		DB:   db,
		Type: "Database",
	}
	rows, err := s.DB.Query(
		"create table if not exists metrics(ID text , MType text, Delta bigint, Value double precision, Hash text)",
	)
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(fmt.Sprintf("Could not establish connection to database %s:", err.Error()))
	}
	return s
}

func (s *DBStorage) Dump() metric.MetricMap {
	metricMap := metric.MetricMap{}
	tx, err := s.DB.Begin()
	if err != nil {
		logging.Warn("Error during transaction creation %s", err.Error())
		return metricMap
	}
	err = FetchRecords(tx, metricMap)
	if err != nil {
		logging.Warn("Error during dumping Database content %s", err.Error())
	}
	return metricMap
}

func (s *DBStorage) Get(metricName string) (metric.Metrics, bool) {
	m := metric.Metrics{ID: metricName}
	tx, err := s.DB.Begin()
	if err != nil {
		logging.Warn("Error during transaction creation %s", err.Error())
		return m, false
	}
	fetchedRecord, err := FetchRecord(tx, m)
	if err != nil {
		logging.Debug("Records were not found")
		return m, false
	}
	return fetchedRecord, true
}

func (s *DBStorage) Update(metrics metric.Metrics) error {
	tx, err := s.DB.Begin()
	if err != nil {
		logging.Warn("Error during transaction creation %s", err.Error())
		return err
	}
	fetchedRecord, err := FetchRecord(tx, metrics)
	if errors.Is(err, sql.ErrNoRows) {
		err = InsertRecord(tx, metrics)
		if err != nil {
			logging.Warn(err.Error())
			return err
		}
	} else if err == nil {
		if metrics.MType == "counter" {
			*metrics.Delta += *fetchedRecord.Delta
		}
		err = UpdateRecord(tx, metrics)
		if err != nil {
			logging.Debug("Problem during transaction process %s", err.Error())
			return err
		}
	} else {
		logging.Debug("Problem during select %s", err.Error())
		return err
	}
	if err := tx.Commit(); err != nil {
		logging.Crit("update drivers: unable to commit: %v", err)
		return err
	}
	return nil

}

func (s *DBStorage) BulkUpdate(aggregatedMetrics map[string]metric.Metrics) error {
	currentMetricMap := metric.MetricMap{}
	tx, err := s.DB.Begin()
	if err != nil {
		logging.Warn("Could not connect to open transaction: %s", err.Error())
		return err
	}
	err = FetchRecords(tx, currentMetricMap)
	if err != nil {
		logging.Warn("Problem during transaction process %s", err.Error())
		return err
	}
	var metricsToInsert []metric.Metrics
	var metricsToUpdate []metric.Metrics
	for id, m := range aggregatedMetrics {
		if cm, found := currentMetricMap[id]; found {
			if m.MType == "counter" {
				*m.Delta += *cm.Delta
			}
			metricsToUpdate = append(metricsToUpdate, m)
		} else {
			metricsToInsert = append(metricsToInsert, m)
		}
	}
	err = InsertRecords(tx, metricsToInsert)
	if err != nil {
		logging.Warn("Could not insert slice of metrics: %s", err.Error())
		return err
	}
	err = UpdateRecords(tx, metricsToUpdate)
	if err != nil {
		logging.Debug("Could not update slice of metrics %s", err.Error())
		return err
	}
	if err := tx.Commit(); err != nil {
		logging.Crit("update drivers: unable to commit: %v", err)
		return err
	}
	return nil
}

func (s *DBStorage) IsAlive() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	logging.Debug("Checking connection to database")
	if err := s.DB.PingContext(ctx); err != nil {
		logging.Warn("Failed: %s", err.Error())
		return err
	}
	logging.Debug("Succeed")
	return nil
}
