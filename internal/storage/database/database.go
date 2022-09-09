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
	s.DB.Query("create table if not exists metrics(ID text , MType text, Delta int, Value double precision, Hash text)")
	return s
}

func (s *DBStorage) Dump() metric.MetricMap {
	metricMap := metric.MetricMap{}
	rows, err := s.DB.Query("select * from metrics;")
	if err != nil {
		fmt.Println("Error During scanning DB")
		fmt.Println(err.Error())
		return metricMap
	}
	defer rows.Close()
	var it int64 = 1
	fmt.Println(rows)
	for rows.Next() {
		m := metric.Metrics{}
		err := rows.Scan(&m.ID, &m.MType, &m.Delta, &m.Value, &m.Hash)
		if err != nil {
			fmt.Println("Error duiring scanning of dumped DB")
			return metricMap
		} else {
			metricMap[m.ID] = m
		}
		fmt.Printf("Finished iteration %d\n", it)
		it += 1
	}
	fmt.Println(metricMap)
	return metricMap
}

func (s *DBStorage) Get(metricName string) (metric.Metrics, bool) {
	m := metric.Metrics{}
	row := s.DB.QueryRowContext(context.Background(),
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
