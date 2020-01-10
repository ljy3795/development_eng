package mylibs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mySQL package
	"log"
	"os"
)

// DB is Database Connection instance
var DB *sql.DB

func init() {
	DB = DbConn()
}

// DbConn returns the Database connection to MariaDB
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MARIADB_USER")
	dbPass := os.Getenv("MARIADB_PASS")
	dbName := "comm"
	dbAddr := "127.0.0.1"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)

	return db
}

// ReplaceQuery returns the number of inserted rows by a query
func ReplaceQuery(db *sql.DB, rows []ROW) int64 {
	var insertedRows int64

	for i := 0; i < len(rows); i++ {
		tempResult := rows[i]
		queryResults, err := db.Exec("REPLACE INTO pm2_daily VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
			tempResult.DT,     // 일자
			tempResult.REGION, // 지역
			tempResult.NO2,
			tempResult.O3,
			tempResult.CO,
			tempResult.SO2,
			tempResult.PM10,
			tempResult.PM25)
		if err != nil {
			log.Fatal(err)
		}

		resultCounts, _ := queryResults.RowsAffected()
		insertedRows += resultCounts
	}
	return insertedRows
}

// DeleteRow do delete SQL operation by primary kyes(dt / region)
func DeleteRow(db *sql.DB, dt string, region string) error {
	tableName := "pm2_daily"
	query := fmt.Sprintf("DELETE FROM %s WHERE (dt = '%s' AND region = '%s')", tableName, dt, region)

	queryResults, err := db.Query(query)

	if err != nil {
		return err
	}

	queryResults.Close()

	return nil
}

// SelectRows returns ROW struct retrieved by SQL SELECT
func SelectRows(db *sql.DB, qry string) ([]*ROW, error) {
	results := make([]*ROW, 0)

	queryResults, err := db.Query(qry)
	if err != nil {
		return nil, err
	}

	for queryResults.Next() {
		tempResults := new(ROW)
		err = queryResults.Scan(&tempResults.DT, &tempResults.REGION, &tempResults.NO2, &tempResults.O3, &tempResults.CO, &tempResults.SO2, &tempResults.PM10, &tempResults.PM25)
		if err != nil {
			return nil, err
		}
		results = append(results, tempResults)
	}

	queryResults.Close()

	return results, nil
}
