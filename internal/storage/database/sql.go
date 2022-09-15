package database

import (
	"database/sql"
	"fmt"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/logging"
	"github.com/aligang/YandexPracticumGoAdvanced/internal/metric"
	"strconv"
)

func ConstructSelectQuery() string {
	return "SELECT ID,MType,Delta,Value,Hash FROM metrics WHERE ID = $1"
}

func FetchRecord(tx *sql.Tx, metrics metric.Metrics) (metric.Metrics, error) {
	fetchedMetrics := metric.Metrics{}
	query := ConstructSelectQuery()
	selectStatement, err := tx.Prepare(query)
	if err != nil {
		logging.Warn("Could not create select statement: %s", err.Error())
		return fetchedMetrics, err
	}
	row := selectStatement.QueryRow(metrics.ID)
	if err := row.Err(); err != nil {
		logging.Warn(err.Error())
	}
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
	if err := rows.Err(); err != nil {
		logging.Warn(err.Error())
		return err
	}
	if err != nil {
		logging.Warn("Error During DB interactionL %s", err.Error())
		return err
	}
	defer rows.Close()
	for rows.Next() {
		m := metric.Metrics{}
		err = rows.Scan(&m.ID, &m.MType, &m.Delta, &m.Value, &m.Hash)
		if err != nil {
			logging.Warn("Error during scanning of dumped DB")
			return err
		}
		metricMap[m.ID] = m
	}
	return nil
}

func ConstructInsertQuery(metrics metric.Metrics) (string, []any) {
	var args []any
	query := ""
	if metrics.MType == "gauge" {
		value := strconv.FormatFloat(*metrics.Value, 'f', -1, 64)
		query = "INSERT INTO metrics (ID, MType, Value, Hash) VALUES($1, $2, $3, $4)"
		args = append(args, metrics.ID, metrics.MType, value, metrics.Hash)
	} else if metrics.MType == "counter" {
		query = "INSERT INTO metrics (ID, MType, Delta, Hash) VALUES($1, $2,  $3, $4)"
		args = append(args, metrics.ID, metrics.MType, *metrics.Delta, metrics.Hash)
	}
	return query, args
}

func InsertRecord(tx *sql.Tx, metrics metric.Metrics) error {
	logging.Debug("Creating New Record")
	insertQuery, args := ConstructInsertQuery(metrics)
	logging.Debug(insertQuery)
	logging.Debug(fmt.Sprintln(args))
	insertStatement, err := tx.Prepare(insertQuery)
	if err != nil {
		logging.Warn("Error during statement preparation %s", err.Error())
		return err
	}
	_, err = insertStatement.Exec(args...)
	if err != nil {
		logging.Warn("Problem during request Execution %s", err.Error())
		if err = tx.Rollback(); err != nil {
			logging.Crit("insert drivers: unable to rollback: %s", err.Error())
		}
		return err
	}
	return nil
}

func InsertRecords(tx *sql.Tx, metricSlice []metric.Metrics) error {
	for _, metric := range metricSlice {
		insertQuery, args := ConstructInsertQuery(metric)
		logging.Debug("Preparing request to DB server: %s", insertQuery)
		insertStatement, err := tx.Prepare(insertQuery)
		if err != nil {
			logging.Warn("Error during statement preparation %s", err.Error())
			return err
		}
		logging.Debug("Executing request to DB server")
		_, err = insertStatement.Exec(args)
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

func ConstructUpdateQuery(metrics metric.Metrics) (string, []any) {
	query := ""
	var args []any
	if metrics.MType == "gauge" {
		value := strconv.FormatFloat(*metrics.Value, 'f', -1, 64)
		query = "UPDATE metrics SET Mtype = $1, Value = $2, Hash = $3 where id = $4"
		args = append(args, metrics.MType, value, metrics.Hash, metrics.ID)
	} else if metrics.MType == "counter" {
		query = "UPDATE metrics SET Mtype = $1, Delta = $2, Hash = $3 where id = $4"
		args = append(args, metrics.MType, *metrics.Delta, metrics.Hash, metrics.ID)
	}
	return query, args
}

func UpdateRecord(tx *sql.Tx, metrics metric.Metrics) error {
	logging.Debug("Updating Old Record")
	updateQuery, args := ConstructUpdateQuery(metrics)
	logging.Debug("Preparing request to DB server: %s", updateQuery)
	logging.Debug(fmt.Sprintln(args))
	updateStatement, err := tx.Prepare(updateQuery)
	if err != nil {
		logging.Warn("Problem during query preparation: %s", err.Error())
		return err
	}
	_, err = updateStatement.Exec(args...)
	if err != nil {
		logging.Warn("Problem during request execution: %s", err.Error())
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
		updateQuery, args := ConstructUpdateQuery(metric)
		logging.Debug("Preparing request to DB server: %s", updateQuery)
		updateStatement, err := tx.Prepare(updateQuery)
		if err != nil {
			logging.Debug("Problem during query preparation %s", err.Error())
		}
		logging.Debug("Executing request")
		_, err = updateStatement.Exec(args)
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
