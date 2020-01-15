package models

// AirQualityDaily is the GORM model definition for Table for air_quality_daily
type AirQualityDaily struct {
	DT     string  `gorm:"PRIMARY_KEY"` // json을 붙여서 재활용
	REGION string  `gorm:"PRIMARY_KEY"`
	NO2    float64 `gorm:"default:0"`
	O3     float64 `gorm:"default:0"`
	CO     float64 `gorm:"default:0"`
	SO2    float64 `gorm:"default:0"`
	PM10   float64 `gorm:"default:0"`
	PM25   float64 `gorm:"default:0"`
}
