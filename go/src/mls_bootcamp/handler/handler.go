package handler

import (
	"fmt"
	"mls_bootcamp/mylibs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ViewHandler rederes view page retrieved by RDS SELECT query results
func ViewHandler(c *gin.Context) {
	fmt.Println("View")

	dt := c.Param("dt")
	region := c.Param("region")

	res, err := mylibs.SelectRows(mylibs.DB, fmt.Sprintf("SELECT * FROM pm2_daily where dt = '%s' AND region = '%s'", dt, region))
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error()))
	} else if len(res) == 0 {
		c.HTML(http.StatusOK, "alert.html", gin.H{
			"org_dt": dt,
			"region": region})
	} else {
		c.HTML(http.StatusOK, "view.html", gin.H{
			"dt":     dt[:4] + "/" + dt[4:6] + "/" + dt[6:8],
			"org_dt": dt,
			"region": region,
			"no2":    res[0].NO2,
			"o3":     res[0].O3,
			"co":     res[0].CO,
			"so2":    res[0].SO2,
			"pm10":   res[0].PM10,
			"pm25":   res[0].PM25,
		})
	}
	fmt.Println("View End")

}

// EditHandler renders retrieved RDS result and send form Data to SaveHandler
func EditHandler(c *gin.Context) {
	fmt.Println("Edit")

	dt := c.Param("dt")
	region := c.Param("region")

	res, err := mylibs.SelectRows(mylibs.DB, fmt.Sprintf("SELECT * FROM pm2_daily where dt = '%s' AND region = '%s'", dt, region))
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error()))
	} else if len(res) == 0 {
		res = make([]*mylibs.ROW, 0)
		res = append(res, new(mylibs.ROW))
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"dt":     dt[:4] + "/" + dt[4:6] + "/" + dt[6:8],
		"org_dt": dt,
		"region": region,
		"no2":    res[0].NO2,
		"o3":     res[0].O3,
		"co":     res[0].CO,
		"so2":    res[0].SO2,
		"pm10":   res[0].PM10,
		"pm25":   res[0].PM25,
	})
	fmt.Println("Edit End")

}

// SaveHandler save retrieved form value from EditHandler to RDS
func SaveHandler(c *gin.Context) {
	fmt.Println("Save")

	// get data from POST form
	originalDate := c.Param("dt")
	region := c.PostForm("region")
	no2, _ := strconv.ParseFloat(c.PostForm("no2"), 64)
	o3, _ := strconv.ParseFloat(c.PostForm("o3"), 64)
	co, _ := strconv.ParseFloat(c.PostForm("co"), 64)
	so2, _ := strconv.ParseFloat(c.PostForm("so2"), 64)
	pm10, _ := strconv.ParseFloat(c.PostForm("pm10"), 64)
	pm25, _ := strconv.ParseFloat(c.PostForm("pm25"), 64)

	// make struct to edit data
	row := mylibs.ROW{originalDate, region, no2, o3, co, so2, pm10, pm25}

	res := make([]mylibs.ROW, 0)
	res = append(res, row)
	mylibs.ReplaceQuery(mylibs.DB, res)
	fmt.Println("Save End")

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/view/%s/%s", originalDate, region))
}

// DeleteHandler delete a row selected by two primary keys (dt / region)
func DeleteHandler(c *gin.Context) {
	fmt.Println("Delete")

	dt := c.Param("dt")
	region := c.Param("region")

	err := mylibs.DeleteRow(mylibs.DB, dt, region)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("Error %s", err.Error()))
	}
	fmt.Println("Delete End")
	c.HTML(http.StatusOK, "alert.html", gin.H{
		"org_dt": dt,
		"region": region,
	})

	// c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/save/%s/%s", dt, region))
}
