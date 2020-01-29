package main

import (
	"io"
	"log"
	"os"
	"time"

	"mls_bootcamp/controller"
	"mls_bootcamp/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var myLogger *log.Logger

func main() {
	myLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

	// 0) Backfill
	backfillStartDate := time.Now().AddDate(0, 0, -1)
	backfillCounts := 10
	utils.BackfillAPICall(backfillStartDate, backfillCounts)

	// 1) Scheduler for daily API Call (everyday on 9am)
	cronInput := "0 8 * * *"
	utils.DailyCronAPI(cronInput)

	// 2) Web Handler
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/api/read/:dt/:region", controller.ReadHandler) // view/20200101/강남구
	r.POST("/api/delete/:dt/:region", controller.DeleteHandler)
	r.POST("/api/create", controller.CreateHandler)

	r.Run(":8080")
}
