package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mls_bootcamp/clients"
	"mls_bootcamp/models"
	"net/http"
	"os"
	"time"
)

// func (row *ROW) getQueryStatement() {
// 	return fmt.Sprintf("")
// }

// GetFromAPI returns API result (APICall resultCount, resultStatus, resultDetail)
func GetFromAPI(startPage int, endPage int, apiCallDate time.Time) (int, models.STATUS, []models.ROW, error) {
	apiCallDateString := apiCallDate.Format("20060102")

	apiKey := os.Getenv("SEOUL_API_KEY")
	urlString := fmt.Sprintf("http://openAPI.seoul.go.kr:8088/%s/json/DailyAverageAirQuality/%d/%d/%s", apiKey, startPage, endPage, apiCallDateString)

	results, err := http.Get(urlString)
	if err != nil {
		return 0, models.STATUS{}, nil, err
	}

	byteData, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return 0, models.STATUS{}, nil, err
	}

	var pm2 models.PM2
	if err := json.Unmarshal(byteData, &pm2); err != nil {
		return 0, models.STATUS{}, nil, err
	}

	resultCount := pm2.PM2.CALLRESULTCOUNT
	resultStatus := pm2.PM2.STATUS // API 호출 결과
	resultDetail := pm2.PM2.ROW    // API 호출 rows

	if resultStatus.CODE != "INFO-000" {
		errorMessage := fmt.Sprintf("API Call status FAIL : CODE = '%s' / MESSAGE = '%s')\n", resultStatus.CODE, resultStatus.MESSAGE)
		log.Fatal(errorMessage)
		return 0, models.STATUS{}, nil, errors.New(errorMessage)
	}

	return resultCount, resultStatus, resultDetail, nil
}

// InsertAPIRows creates new intserted rows retreived by API Call
func InsertAPIRows(rows []models.ROW) int64 {
	var insertedRows int64

	for i := 0; i < len(rows); i++ {
		tempResult := models.AirQualityDaily(rows[i])

		err := clients.CreateNewRow(tempResult)
		if err != nil {
			clients.UpdateRow(tempResult)
		}
		insertedRows++
	}
	return insertedRows
}
