package services

import (
	"fmt"
	"log"
	"os"
	"time"
)

var myLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

// DailyAPICall retrieves Daily data through GetFromAPI and Insert to RDS usning InsertAPIRows
func DailyAPICall() {
	apiCallDate := time.Now()

	startPage := 1
	endPage := 1000

	_, _, rows, err := GetFromAPI(startPage, endPage, apiCallDate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 2) DB에 적재
	insertedRows := InsertAPIRows(rows)
	myLogger.Printf("%d row inserted.\n", insertedRows)
	myLogger.Println(apiCallDate.Format("20060102"), " is done")
}

// BackfillAPICall is function for do backfill from backfillStartDate to the number of backfillCounts days
func BackfillAPICall(backfillStartDate time.Time, backfillCounts int) {
	startPage := 1
	endPage := 1000

	for i := 0; i < backfillCounts; i++ {

		_, _, rows, err := GetFromAPI(startPage, endPage, backfillStartDate)
		if err != nil {
			fmt.Println(err.Error())
		}

		insertedRows := InsertAPIRows(rows)
		myLogger.Printf("%d row inserted.\n", insertedRows)
		myLogger.Println(backfillStartDate.Format("20060102"), " is done")
		backfillStartDate = backfillStartDate.AddDate(0, 0, -1)
	}
}
