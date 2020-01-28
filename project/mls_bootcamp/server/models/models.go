package models

// PM2 is a main struct conatining the result retured by API Call
type PM2 struct {
	PM2 PM2Body `json:"DailyAverageAirQuality"`
}

// PM2Body is a body struct containing three API Call resutls(CALLRESULTCOUNT / STATUS / ROW)
type PM2Body struct {
	CallResultCount int               `json:"list_total_count"`
	Status          Status            `json:"RESULT"`
	AirQualityDaily []AirQualityDaily `json:"row"`
}

// Status is a struct containing returned API call status (CODE & MESSAGE)
type Status struct {
	Code    string `json:"CODE"`
	Message string `json:"MESSAGE"`
}

// AirQualityDaily is a struct cotaining PM2 information retreived from API
type AirQualityDaily struct {
	DT     string  `json:"MSRDT_DE" gorm:"PRIMARY_KEY"`
	Region string  `json:"MSRSTE_NM" gorm:"PRIMARY_KEY"`
	NO2    float64 `json:"NO2" gorm:"default:0"`
	O3     float64 `json:"O3" gorm:"default:0"`
	CO     float64 `json:"CO" gorm:"default:0"`
	SO2    float64 `json:"SO2" gorm:"default:0"`
	PM10   float64 `json:"PM10" gorm:"default:0"`
	PM25   float64 `json:"PM25" gorm:"default:0"`
}
