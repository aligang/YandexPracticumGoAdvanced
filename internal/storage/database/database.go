package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"log"
	"time"
)
import _ "github.com/jackc/pgx/v4/stdlib"

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
	s.DB.Query("create table if not exists metrics(ID text , MType text, Delta bigint, Value double precision, Hash text)")
	return s
}

func (s *DBStorage) Dump() metric.MetricMap {
	metricMap := metric.MetricMap{}
	err := FetchRecords(s.DB, metricMap)
	if err != nil {
		fmt.Println("Error during dumping Database content")
		fmt.Println(err.Error())
	}
	return metricMap
}

func (s *DBStorage) Get(metricName string) (metric.Metrics, bool) {
	m := metric.Metrics{}
	row := s.DB.QueryRow(
		fmt.Sprintf("select ID,MType,Delta,Value,Hash from metrics where ID = '%s'", metricName),
	)
	err := row.Scan(&m.ID, &m.MType, &m.Delta, &m.Value, &m.Hash)
	if err != nil {
		fmt.Println("Records were not found")
		return m, false
	}
	fmt.Println("Record were found")
	return m, true
}

func (s *DBStorage) Update(metrics metric.Metrics) {
	tx, err := s.DB.Begin()
	fetchedRecord, err := FetchRecord(tx, metrics)
	if err == sql.ErrNoRows {
		err = InsertRecord(tx, metrics)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else if err == nil {
		if metrics.MType == "counter" {
			*metrics.Delta += *fetchedRecord.Delta
		}
		err = UpdateRecord(tx, metrics)
	} else {
		fmt.Println("Problem during select")
		fmt.Println(err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("update drivers: unable to commit: %v", err)
		return
	}

}

func (s *DBStorage) BulkUpdate(metrics []metric.Metrics) {
	currentMetrics := s.Dump()
	var metricsToInsert []metric.Metrics
	var metricsToUpdate []metric.Metrics
	var aggregatedMetrics map[string]metric.Metrics

	for _, m := range metrics {
		_, found := aggregatedMetrics[m.ID]
		if m.MType == "counter" && found {
			*aggregatedMetrics[m.ID].Delta = +*m.Delta
		} else {
			aggregatedMetrics[m.ID] = m
		}
	}
	for id, m := range aggregatedMetrics {
		if cm, found := currentMetrics[id]; found {
			if m.MType == "counter" {
				*m.Delta = +*cm.Delta
			}
			metricsToUpdate = append(metricsToUpdate, m)
		} else {
			metricsToInsert = append(metricsToInsert, m)
		}
	}
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println("Could not connect to open transaction")
		fmt.Println(err.Error())
		return
	}

	err = InsertRecords(tx, metrics)
	if err != nil {
		fmt.Println("Could not insert slice of metrics")
		fmt.Println(err.Error())
	}
	err = UpdateRecords(tx, metrics)
	if err != nil {
		fmt.Println("Could not update slice of metrics")
		fmt.Println(err.Error())
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("update drivers: unable to commit: %v", err)
		return
	}
	fmt.Println("Succsefully put")
}

func (s *DBStorage) IsAlive() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Printf("Checking connection to database\n")
	if err := s.DB.PingContext(ctx); err != nil {
		fmt.Println("Failed")
		fmt.Println(err)
		return err
	}
	return nil
}
