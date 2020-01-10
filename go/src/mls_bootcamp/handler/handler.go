package handler

import (
	"fmt"
	"mls_bootcamp/mylibs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler
func ViewHandler(c *gin.Context) {
	dt := c.Param("dt")
	region := c.Param("region")

	db := mylibs.DbConn()
	defer db.Close()

	res := mylibs.SelectRows(db, fmt.Sprintf("SELECT * FROM pm2_daily where dt = '%s' AND region = '%s'", dt, region))
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

func EditHandler(c *gin.Context) {
	dt := c.Param("dt")
	region := c.Param("region")

	db := mylibs.DbConn()
	defer db.Close()

	res := mylibs.SelectRows(db, fmt.Sprintf("SELECT * FROM pm2_daily where dt = '%s' AND region = '%s'", dt, region))
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
}

// Save Handler
func SaveHandler(c *gin.Context) {
	// get data from POST form
	org_dt := c.Param("dt")
	region := c.PostForm("region")
	no2, _ := strconv.ParseFloat(c.PostForm("no2"), 64)
	o3, _ := strconv.ParseFloat(c.PostForm("o3"), 64)
	co, _ := strconv.ParseFloat(c.PostForm("co"), 64)
	so2, _ := strconv.ParseFloat(c.PostForm("so2"), 64)
	pm10, _ := strconv.ParseFloat(c.PostForm("pm10"), 64)
	pm25, _ := strconv.ParseFloat(c.PostForm("pm25"), 64)

	// make struct to edit data
	row := mylibs.ROW{org_dt, region, no2, o3, co, so2, pm10, pm25}

	db := mylibs.DbConn()
	defer db.Close()

	res := make([]mylibs.ROW, 0)
	res = append(res, row)
	mylibs.ReplaceQuery(db, res)

	c.Redirect(http.StatusMovedPermanently, "/view/"+org_dt+"/"+region)
}
