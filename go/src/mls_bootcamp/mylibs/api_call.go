package mylibs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// PM2 is a main struct conatining the result retured by API Call
type PM2 struct {
	PM2 PM2Body `json:"DailyAverageAirQuality"`
}

// PM2Body is a body struct containing three API Call resutls(CALLRESULTCOUNT / STATUS / ROW)
type PM2Body struct {
	CALLRESULTCOUNT int    `json:"list_total_count"`
	STATUS          STATUS `json:"RESULT"`
	ROW             []ROW  `json:"row"`
}

// STATUS is a struct containing returned API call status (CODE & MESSAGE)
type STATUS struct {
	CODE    string `json:"CODE"`
	MESSAGE string `json:"MESSAGE"`
}

// ROW is a struct cotaining PM2 information retreived from API
type ROW struct {
	DT     string  `json:"MSRDT_DE"`
	REGION string  `json:"MSRSTE_NM"`
	NO2    float64 `json:"NO2"`
	O3     float64 `json:"O3"`
	CO     float64 `json:"CO"`
	SO2    float64 `json:"SO2"`
	PM10   float64 `json:"PM10"`
	PM25   float64 `json:"PM25"`
}

// func (row *ROW) getQueryStatement() {
// 	return fmt.Sprintf("")
// }

// GetFromAPI returns API result (APICall resultCount, resultStatus, resultDetail)
func GetFromAPI(startPage int, endPage int, apiCallDate time.Time) (int, STATUS, []ROW, error) {
	apiCallDateString := apiCallDate.Format("20060102")

	apiKey := os.Getenv("SEOUL_API_KEY")
	urlString := fmt.Sprintf("http://openAPI.seoul.go.kr:8088/%s/json/DailyAverageAirQuality/%d/%d/%s", apiKey, startPage, endPage, apiCallDateString)

	results, err := http.Get(urlString)
	if err != nil {
		return 0, STATUS{}, nil, err
	}

	byteData, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return 0, STATUS{}, nil, err
	}

	var pm2 PM2
	if err := json.Unmarshal(byteData, &pm2); err != nil {
		return 0, STATUS{}, nil, err
	}

	resultCount := pm2.PM2.CALLRESULTCOUNT
	resultStatus := pm2.PM2.STATUS // API 호출 결과
	resultDetail := pm2.PM2.ROW    // API 호출 rows

	if resultStatus.CODE != "INFO-000" {
		errorMessage := fmt.Sprintf("API Call status FAIL : CODE = '%s' / MESSAGE = '%s')\n", resultStatus.CODE, resultStatus.MESSAGE)
		log.Fatal(errorMessage)
		return 0, STATUS{}, nil, errors.New(errorMessage)
	}

	return resultCount, resultStatus, resultDetail, nil
}
