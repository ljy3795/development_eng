package models

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
