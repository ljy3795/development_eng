package controller

import (
	"fmt"
	"mls_bootcamp/clients"
	"mls_bootcamp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ViewHandler rederes view page retrieved by RDS SELECT query results
func ViewHandler(c *gin.Context) {
	fmt.Println("View")

	dt := c.Param("dt")
	region := c.Param("region")

	res, err := clients.SelectRow(dt, region)

	if err != nil && err.Error() != "record not found" {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error())) // c.Json
	} else if res == (models.AirQualityDaily{}) {
		c.HTML(http.StatusOK, "alert.html", gin.H{
			"message": "No Data!",
			"org_dt":  dt,
			"region":  region})
	} else {
		c.HTML(http.StatusOK, "view.html", gin.H{
			"dt":     dt[:4] + "/" + dt[4:6] + "/" + dt[6:8],
			"org_dt": dt,
			"region": region,
			"no2":    res.NO2,
			"o3":     res.O3,
			"co":     res.CO,
			"so2":    res.SO2,
			"pm10":   res.PM10,
			"pm25":   res.PM25,
		})
	}
	fmt.Println("View End")

}

// EditHandler renders retrieved RDS result and send form Data to SaveHandler
func EditHandler(c *gin.Context) {
	fmt.Println("Edit")

	dt := c.Param("dt")
	region := c.Param("region")

	res, err := clients.SelectRow(dt, region)
	if err != nil && err.Error() != "record not found" {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error()))
	}

	fmt.Println(dt)
	fmt.Println(region)

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"dt":     dt[:4] + "/" + dt[4:6] + "/" + dt[6:8],
		"org_dt": dt,
		"region": region,
		"no2":    res.NO2,
		"o3":     res.O3,
		"co":     res.CO,
		"so2":    res.SO2,
		"pm10":   res.PM10,
		"pm25":   res.PM25,
		"type":   "Edit",
	})
	fmt.Println("Edit End")

}

// SaveHandler save retrieved form value from EditHandler to RDS
func SaveHandler(c *gin.Context) {
	fmt.Println("Save")

	// get data from POST form
	dt := c.Param("dt")
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

	if err != nil && err.Error() == "record not found" {
		clients.CreateNewRow(row)
		fmt.Println("Save End")
	} else {
		err := clients.UpdateRow(row)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Save End")
		}
	}
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/view/%s/%s", dt, region))
}

// DeleteHandler delete a row selected by two primary keys (dt / region)
func DeleteHandler(c *gin.Context) {
	fmt.Println("Delete")

	dt := c.Param("dt")
	region := c.Param("region")

	err := clients.DeleteRow(dt, region)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error()))
	}
	fmt.Println("Delete End")
	c.HTML(http.StatusOK, "alert.html", gin.H{
		"message": "Successfully Removed",
		"org_dt":  dt,
		"region":  region,
	})
}

// AddHandler renders retrieved RDS result and send form Data to SaveHandler
func AddHandler(c *gin.Context) {
	fmt.Println("Add")

	baseRow := models.AirQualityDaily{}

	c.HTML(http.StatusOK, "add.html", gin.H{
		"dt":     "",
		"org_dt": "",
		"region": "",
		"no2":    baseRow.NO2,
		"o3":     baseRow.O3,
		"co":     baseRow.CO,
		"so2":    baseRow.SO2,
		"pm10":   baseRow.PM10,
		"pm25":   baseRow.PM25,
		"type":   "Add",
	})
	fmt.Println("Add End")

}

// AddSaveHandler aa
func AddSaveHandler(c *gin.Context) {
	fmt.Println("TempSave")

	dt := c.PostForm("dt")
	region := c.PostForm("region")
	no2, _ := strconv.ParseFloat(c.PostForm("no2"), 64)
	o3, _ := strconv.ParseFloat(c.PostForm("o3"), 64)
	co, _ := strconv.ParseFloat(c.PostForm("co"), 64)
	so2, _ := strconv.ParseFloat(c.PostForm("so2"), 64)
	pm10, _ := strconv.ParseFloat(c.PostForm("pm10"), 64)
	pm25, _ := strconv.ParseFloat(c.PostForm("pm25"), 64)

	// make struct to edit data
	row := models.AirQualityDaily{DT: dt, REGION: region, NO2: no2, O3: o3, CO: co, SO2: so2, PM10: pm10, PM25: pm25}
	fmt.Println(row)

	clients.CreateNewRow(row)
	fmt.Println("TempSave End")

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/view/%s/%s", dt, region))

}
