package mylibs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PM2 struct {
	PM2 PM2Body `json:"DailyAverageAirQuality"`
}

type PM2Body struct {
	List_total_count int    `json:"list_total_count"`
	RESULT           RESULT `json:"RESULT"`
	ROW              []ROW  `json:"row"`
}

type RESULT struct {
	CODE    string `json:"CODE"`
	MESSAGE string `json:"MESSAGE"`
}

type ROW struct {
	MSRDTDE  string  `json:"MSRDT_DE"`
	MSRSTENM string  `json:"MSRSTE_NM"`
	NO2      float64 `json:"NO2"`
	O3       float64 `json:"O3"`
	CO       float64 `json:"CO"`
	SO2      float64 `json:"SO2"`
	PM10     float64 `json:"PM10"`
	PM25     float64 `json:"PM25"`
}

// 서울 Open API Call
func GetFromAPI(sta_i int, end_i int, strd_dt time.Time) (int, RESULT, []ROW) {
	strd_dt_str := strd_dt.Format("20060102")

	api_key := "6f536d76626c6a793534526f6f7a57"
	url_txt := fmt.Sprintf("http://openAPI.seoul.go.kr:8088/%s/json/DailyAverageAirQuality/%d/%d/%s", api_key, sta_i, end_i, strd_dt_str)

	res, err := http.Get(url_txt)
	if err != nil {
		panic(err)
	}
	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pm2 PM2
	err = json.Unmarshal(byteData, &pm2)
	if err != nil {
		panic(err)
	}

	// var aa map[string]interface{}
	// err = json.Unmarshal(byteData, &aa)

	r1 := pm2.PM2.List_total_count // API 카운트 수
	r2 := pm2.PM2.RESULT           // API 호출 결과
	r3 := pm2.PM2.ROW              // API 호출 rows

	// check API data is okay
	if r2.CODE != "INFO-000" {
		panic("The API data is ERROR, " + r2.CODE + " " + r2.MESSAGE)
	}

	return r1, r2, r3
}
