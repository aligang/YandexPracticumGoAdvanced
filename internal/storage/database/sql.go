package database

import (
	"database/sql"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"log"
	"strconv"
)

func ConstructSelectQuery(metrics metric.Metrics) string {
	query := fmt.Sprintf("select ID,MType,Delta,Value,Hash from metrics where ID = '%s'", metrics.ID)
	return query
}

func FetchRecord(tx *sql.Tx, metrics metric.Metrics) (metric.Metrics, error) {
	fetchedMetrics := metric.Metrics{}
	query := ConstructSelectQuery(metrics)
	selectStatement, err := tx.Prepare(query)
	if err != nil {
		fmt.Println("Could not create select statement")
		fmt.Println(err.Error())
		return fetchedMetrics, err
	}
	row := selectStatement.QueryRow()
	err = row.Scan(&fetchedMetrics.ID, &fetchedMetrics.MType, &fetchedMetrics.Delta, &fetchedMetrics.Value, &fetchedMetrics.Hash)
	if err != nil {
		fmt.Println("Could not decode Database Server response")
		fmt.Println(err.Error())
		return fetchedMetrics, err
	}
	return fetchedMetrics, nil
}

func FetchRecords(db *sql.DB, metricMap metric.MetricMap) error {
	rows, err := db.Query("select * from metrics;")
	if err != nil {
		fmt.Println("Error During scanning DB")
		fmt.Println(err.Error())
		return err
	}
	defer rows.Close()
	var it int64 = 1
	fmt.Println(rows)
	for rows.Next() {
		m := metric.Metrics{}
		err := rows.Scan(&m.ID, &m.MType, &m.Delta, &m.Value, &m.Hash)
		if err != nil {
			fmt.Println("Error duiring scanning of dumped DB")
			return err
		} else {
			metricMap[m.ID] = m
		}
		fmt.Printf("Finished iteration %d\n", it)
		it += 1
	}
	fmt.Println(metricMap)
	return nil
}

func ConstructInsertQuery(metrics metric.Metrics) string {
	query := ""
	if metrics.MType == "gauge" {
		value := strconv.FormatFloat(*metrics.Value, 'f', -1, 64)
		query = fmt.Sprintf("INSERT INTO metrics (ID, MType, Value, Hash) VALUES('%s', '%s',  %s, '%s')",
			metrics.ID, metrics.MType, value, metrics.Hash)
	} else if metrics.MType == "counter" {
		query = fmt.Sprintf("INSERT INTO metrics (ID, MType, Delta, Hash) VALUES('%s', '%s',  %d, '%s')",
			metrics.ID, metrics.MType, *metrics.Delta, metrics.Hash)
	}
	return query
}

func InsertRecord(tx *sql.Tx, metrics metric.Metrics) error {
	fmt.Println("Creating New Record")
	insertQuery := ConstructInsertQuery(metrics)
	fmt.Println(insertQuery)
	insertStatement, err := tx.Prepare(insertQuery)
	if err != nil {
		fmt.Println("Error during statement preparation")
		fmt.Println(err.Error())
		return err
	}
	_, err = insertStatement.Exec()

	if err != nil {
		fmt.Println(err.Error())
		if err = tx.Rollback(); err != nil {

			log.Fatalf("insert drivers: unable to rollback: %v", err)
		}
		return err
	}
	return nil
}

func InsertRecords(tx *sql.Tx, metricSlice []metric.Metrics) error {
	for _, metric := range metricSlice {
		insertQuery := ConstructInsertQuery(metric)
		fmt.Println(insertQuery)
		insertStatement, err := tx.Prepare(insertQuery)
		if err != nil {
			fmt.Println("Error during statement preparation")
			fmt.Println(err.Error())
			return err
		}
		_, err = insertStatement.Exec()
		if err != nil {
			fmt.Println(err.Error())
			if err = tx.Rollback(); err != nil {
				log.Fatalf("insert drivers: unable to rollback: %v", err)
			}
			return err
		}
	}
	return nil
}

func ConstructUpdateQuery(metrics metric.Metrics) string {
	query := ""
	if metrics.MType == "gauge" {
		value := strconv.FormatFloat(*metrics.Value, 'f', -1, 64)
		query = fmt.Sprintf("UPDATE metrics SET Mtype = '%s', Value = '%s', Hash = '%s' where id = '%s'",
			metrics.MType, value, metrics.Hash, metrics.ID)
	} else if metrics.MType == "counter" {
		query = fmt.Sprintf("UPDATE metrics SET Mtype = '%s', Delta = '%d', Hash = '%s' where id = '%s'",
			metrics.MType, *metrics.Delta, metrics.Hash, metrics.ID)
	}
	return query
}

func UpdateRecord(tx *sql.Tx, metrics metric.Metrics) error {
	fmt.Println("Updating Old Record")
	updateQuery := ConstructUpdateQuery(metrics)
	fmt.Println(updateQuery)
	updateStatement, err := tx.Prepare(updateQuery)
	if err != nil {
		fmt.Println("Problem during query preparation")
		fmt.Println(err.Error())
	}
	_, err = updateStatement.Exec()
	if err != nil {
		if err = tx.Rollback(); err != nil {
			log.Fatalf("update drivers: unable to rollback: %v", err)
		}
		return err
	}
	return nil
}

func UpdateRecords(tx *sql.Tx, metrics []metric.Metrics) error {
	fmt.Println("Updating metrics")
	for _, metric := range metrics {
		updateQuery := ConstructUpdateQuery(metric)
		fmt.Println(updateQuery)
		updateStatement, err := tx.Prepare(updateQuery)
		if err != nil {
			fmt.Println("Problem during query preparation")
			fmt.Println(err.Error())
		}
		fmt.Println("Sending...")
		_, err = updateStatement.Exec()
		if err != nil {
			fmt.Println("Trouble")
			if err = tx.Rollback(); err != nil {
				log.Fatalf("update drivers: unable to rollback: %v", err)
			}
			return err
		}
		fmt.Println("No Trouble")
	}
	return nil
}
