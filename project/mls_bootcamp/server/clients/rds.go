package clients

import (
	"errors"
	"fmt"
	"log"
	"mls_bootcamp/models"
	"os"

	_ "github.com/go-sql-driver/mysql" // mySQL package
	"github.com/jinzhu/gorm"
)

// DB is Database Connection instance
var DB *gorm.DB

func init() {
	DB = DbConn()
	if DB.HasTable(&models.AirQualityDaily{}) == false {
		DB.Set("gorm:table_options", "CHARSET=UTF8").AutoMigrate(&models.AirQualityDaily{})
	}
}

// DbConn returns the Database connection to MariaDB
func DbConn() (db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("MARIADB_USER")
	dbPass := os.Getenv("MARIADB_PASS")
	dbAddr := os.Getenv("MARIADB_ADDR")
	dbPort := os.Getenv("MARIADB_PORT")
	dbName := "comm"

	db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(50)

	return db
}

// SelectRow returns SELECT query result by two primary keys(dt / region)
func SelectRow(dt string, region string) (models.AirQualityDaily, error) {
	var res models.AirQualityDaily
	err := DB.Model(&models.AirQualityDaily{}).Where("dt = ? AND region = ?", dt, region).Find(&res).Error

	if err != nil {
		return models.AirQualityDaily{}, err
	}

	return res, nil
}

// CreateNewRow operates adding New Row
func CreateNewRow(row models.AirQualityDaily) error {
	err := DB.Model(&models.AirQualityDaily{}).Create(&row).Error

	if err != nil {
		return err
	}
	return nil
}

// DeleteRow operates Delete query
func DeleteRow(dt string, region string) error {
	if res, _ := SelectRow(dt, region); res == (models.AirQualityDaily{}) {
		return errors.New("No rows to DELETE")
	}

	err := DB.Where("dt = ? And region = ?", dt, region).Delete(&models.AirQualityDaily{}).Error

	if err != nil {
		return err
	}
	return nil
}

// UpdateRow operates Update query
func UpdateRow(newRow models.AirQualityDaily) error {
	err := DB.Model(&models.AirQualityDaily{}).UpdateColumns(&newRow).Error

	if err != nil {
		return err
	}

	return nil
}
