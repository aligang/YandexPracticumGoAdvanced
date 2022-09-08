package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/config"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
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
	return s
}

func (s *DBStorage) Dump() metric.MetricMap {
	return metric.MetricMap{}
}

func (s *DBStorage) Get(metricName string) (metric.Metrics, bool) {
	return metric.Metrics{}, true
}

func (s *DBStorage) Update(metrics metric.Metrics) {

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
