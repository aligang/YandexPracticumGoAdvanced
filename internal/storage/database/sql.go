package database

import (
	"database/sql"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
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
		logging.Warn("Could not create select statement: %s", err.Error())
		return fetchedMetrics, err
	}
	row := selectStatement.QueryRow()
	err = row.Scan(&fetchedMetrics.ID, &fetchedMetrics.MType, &fetchedMetrics.Delta, &fetchedMetrics.Value, &fetchedMetrics.Hash)
	if err != nil {
		logging.Warn("Could not decode Database Server response: %s", err.Error())
		return fetchedMetrics, err
	}
	return fetchedMetrics, nil
}

func FetchRecords(tx *sql.Tx, metricMap metric.MetricMap) error {
	fetchStatement, err := tx.Prepare("select * from metrics;")
	if err != nil {
		logging.Warn("Error Preparing Statement %s", err.Error())
		return err
	}
	rows, err := fetchStatement.Query()
	if err != nil {
		logging.Warn("Error During DB interaction %s": err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		m := metric.Metrics{}
		err := rows.Scan(&m.ID, &m.MType, &m.Delta, &m.Value, &m.Hash)
		if err != nil {
			logging.Warn("Error during scanning of dumped DB")
			return err
		} else {
			metricMap[m.ID] = m
		}
	}
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
	logging.Debug("Creating New Record")
	insertQuery := ConstructInsertQuery(metrics)
	fmt.Println(insertQuery)
	insertStatement, err := tx.Prepare(insertQuery)
	if err != nil {
		logging.Warn("Error during statement preparation %s", err.Error())
		return err
	}
	_, err = insertStatement.Exec()

	if err != nil {
		logging.Warn(err.Error())
		if err = tx.Rollback(); err != nil {

			logging.Crit("insert drivers: unable to rollback: %s", err.Error())
		}
		return err
	}
	return nil
}

func InsertRecords(tx *sql.Tx, metricSlice []metric.Metrics) error {
	for _, metric := range metricSlice {
		insertQuery := ConstructInsertQuery(metric)
		logging.Debug("Preparing request to DB server: %s",insertQuery)
		insertStatement, err := tx.Prepare(insertQuery)
		if err != nil {
			logging.Warn("Error during statement preparation %s", err.Error())
			return err
		}
		logging.Debug("Executing request to DB server")
		_, err = insertStatement.Exec()
		if err != nil {
			fmt.Println(err.Error())
			if err = tx.Rollback(); err != nil {
				logging.Crit("insert drivers: unable to rollback: %s", err.Error())
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
	logging.Debug("Updating Old Record")
	updateQuery := ConstructUpdateQuery(metrics)
	logging.Debug( "Preparing request to DB server: %s", updateQuery)
	updateStatement, err := tx.Prepare(updateQuery)
	if err != nil {
		logging.Warn("Problem during query preparation: %s", err.Error())
	}
	_, err = updateStatement.Exec()
	if err != nil {
		if err = tx.Rollback(); err != nil {
			logging.Crit("update drivers: unable to rollback: %s", err.Error())
		}
		return err
	}
	return nil
}

func UpdateRecords(tx *sql.Tx, metrics []metric.Metrics) error {
	logging.Debug("Updating metrics")
	for _, metric := range metrics {
		updateQuery := ConstructUpdateQuery(metric)
		logging.Debug("Preparing request to DB server: %s", updateQuery)
		updateStatement, err := tx.Prepare(updateQuery)
		if err != nil {
			logging.Debug("Problem during query preparation %s", err.Error())
		}
		logging.Debug("Executing request")
		_, err = updateStatement.Exec()
		if err != nil {
			logging.Warn(err.Error())
			if err = tx.Rollback(); err != nil {
				logging.Crit("update drivers: unable to rollback: %v", err)
			}
			return err
		}
		fmt.Println("Bulk update data push was successful")
	}
	return nil
}
