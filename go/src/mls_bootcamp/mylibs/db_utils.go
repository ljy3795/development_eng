package mylibs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var myLogger *log.Logger

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "colin"
	dbPass := "aa"
	dbName := "comm"
	dbAddr := "127.0.0.1"
	dbPort := "3306"

	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ReplaceQuery(db *sql.DB, rows []ROW) int64 {
	var ins_rows int64
	for i := 0; i < len(rows); i++ {
		tmp_res := rows[i]
		result, err := db.Exec("REPLACE INTO pm2_daily VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
			tmp_res.MSRDTDE,  // 일자
			tmp_res.MSRSTENM, // 지역
			tmp_res.NO2,
			tmp_res.O3,
			tmp_res.CO,
			tmp_res.SO2,
			tmp_res.PM10,
			tmp_res.PM25)
		if err != nil {
			log.Fatal(err)
		}

		up_int, _ := result.RowsAffected()
		ins_rows += up_int
	}
	return ins_rows
}

// delete row by primary key columns (dt, region)
func DeleteRow(db *sql.DB, dt string, region string) {
	table_name := "pm2_daily"
	qry := fmt.Sprintf("DELETE FROM %s WHERE (dt = '%s' AND region = '%s')", table_name, dt, region)
	fmt.Println(qry)

	drop_q, err := db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	drop_q.Close()
}

func SelectRows(db *sql.DB, qry string) []*ROW {
	res := make([]*ROW, 0)

	results, err := db.Query(qry)
	if err != nil {
		panic(err.Error)
	}

	for results.Next() {
		tmp_res := new(ROW)
		err = results.Scan(&tmp_res.MSRDTDE, &tmp_res.MSRSTENM, &tmp_res.NO2, &tmp_res.O3, &tmp_res.CO, &tmp_res.SO2, &tmp_res.PM10, &tmp_res.PM25)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, tmp_res)
	}

	return res
}
