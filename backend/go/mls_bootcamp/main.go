package main

import (
	"io"
	"log"
	"os"
	"time"

	"mls_bootcamp/clients"
	"mls_bootcamp/controller"
	"mls_bootcamp/services"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// Logger
var myLogger *log.Logger

func dailyAPICall() {
	apiCallDate := time.Now()

	startPage := 1
	endPage := 1000

	_, _, rows, err := services.GetFromAPI(startPage, endPage, apiCallDate)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 2) DB에 적재
	insertedRows := services.InsertAPIRows(rows)
	myLogger.Printf("%d row inserted.\n", insertedRows)
	myLogger.Println(apiCallDate.Format("20060102"), " is done")
}

func main() {
	// myLogger = log.New(os.Stdout, "Info: ", log.LstdFlags)
	myLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

	DB := clients.DB
	defer DB.Close()

	// 0) Backfill
	backfillStartDate := time.Now().AddDate(0, 0, -1)
	backfillCounts := 50

	for i := 0; i < backfillCounts; i++ {
		startPage := 1
		endPage := 1000

		_, _, rows, err := services.GetFromAPI(startPage, endPage, backfillStartDate)
		if err != nil {
			log.Fatal(err.Error())
		}

		insertedRows := services.InsertAPIRows(rows)
		myLogger.Printf("%d row inserted.\n", insertedRows)
		myLogger.Println(backfillStartDate.Format("20060102"), " is done")
		backfillStartDate = backfillStartDate.AddDate(0, 0, -1)
	}

	// 1) Scheduler for daily API Call (everyday on 9am)
	c := cron.New()
	c.AddFunc("0 8 * * *", func() {
		dailyAPICall()
	})
	c.Start()

	// 2) Web Handler
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// web
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/view/:dt/:region", controller.ViewHandler) // view/20200101/강남구
	r.GET("/edit/:dt/:region", controller.EditHandler)
	r.POST("/delete/:dt/:region", controller.DeleteHandler)
	r.POST("/save/:dt/:region", controller.SaveHandler)
	r.POST("/add/:dt", controller.AddHandler)
	r.POST("/add_save", controller.AddSaveHandler)

	r.Run(":8080")

	defer c.Stop()
}
