package controller

import (
	"mls_bootcamp/clients"
	"mls_bootcamp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReadHandler rederes view page retrieved by RDS SELECT query results
func ReadHandler(c *gin.Context) {
	dt := c.Param("dt")
	region := c.Param("region")

	res, err := clients.SelectRow(dt, region)

	c.Header("Access-Control-Allow-Origin", "*")
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else if res == (models.AirQualityDaily{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No Data!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"dt":     dt,
			"region": region,
			"no2":    res.NO2,
			"o3":     res.O3,
			"co":     res.CO,
			"so2":    res.SO2,
			"pm10":   res.PM10,
			"pm25":   res.PM25,
		})
	}
}

// DeleteHandler delete a row selected by two primary keys (dt / region)
func DeleteHandler(c *gin.Context) {
	dt := c.Param("dt")
	region := c.Param("region")

	err := clients.DeleteRow(dt, region)

	c.Header("Access-Control-Allow-Origin", "*")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Removed",
		})
	}
}

// CreateHandler save retrieved form value to RDS
func CreateHandler(c *gin.Context) {
	// get data from POST form
	dt := c.PostForm("dt")
	region := c.PostForm("region")

	// Error 받아서 redirect or JS에서 해결
	no2, _ := strconv.ParseFloat(c.PostForm("no2"), 64)
	o3, _ := strconv.ParseFloat(c.PostForm("o3"), 64)
	co, _ := strconv.ParseFloat(c.PostForm("co"), 64)
	so2, _ := strconv.ParseFloat(c.PostForm("so2"), 64)
	pm10, _ := strconv.ParseFloat(c.PostForm("pm10"), 64)
	pm25, _ := strconv.ParseFloat(c.PostForm("pm25"), 64)

	// make struct to edit data
	row := models.AirQualityDaily{DT: dt, REGION: region, NO2: no2, O3: o3, CO: co, SO2: so2, PM10: pm10, PM25: pm25}

	_, err := clients.SelectRow(dt, region)

	c.Header("Access-Control-Allow-Origin", "*")
	if err != nil && err.Error() == "record not found" {
		clients.CreateNewRow(row)
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Added",
		})
	} else {
		err := clients.UpdateRow(row)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "New Row is Successfully Updated",
			})
		}
	}
}
