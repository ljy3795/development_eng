package main

import (
	"io"
	"log"
	"os"
	"time"

	"mls_bootcamp/handler"
	"mls_bootcamp/mylibs"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// Logger
var myLogger *log.Logger

func dailyAPICall() {
	apiCallDate := time.Now()

	startPage := 1
	endPage := 1000

	_, _, rows, err := mylibs.GetFromAPI(startPage, endPage, apiCallDate)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 2) DB에 적재
	insertedRows := mylibs.ReplaceQuery(mylibs.DB, rows)
	myLogger.Printf("%d row inserted.\n", insertedRows)
	myLogger.Println(apiCallDate.Format("20060102"), " is done")
}

func main() {
	// myLogger = log.New(os.Stdout, "Info: ", log.LstdFlags)
	myLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

	DB := mylibs.DB
	defer DB.Close()

	// 0) Backfill
	backfillStartDate := time.Now().AddDate(0, 0, -1)
	backfillCounts := 30

	for i := 0; i < backfillCounts; i++ {
		startPage := 1
		endPage := 1000

		_, _, rows, err := mylibs.GetFromAPI(startPage, endPage, backfillStartDate)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Load to MariaDB
		insertedRows := mylibs.ReplaceQuery(DB, rows)
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
	// sig := make(chan os.Signal)
	// signal.Notify(sig, os.Interrupt, os.Kill)
	// <-sig

	// 2) Web Handler
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// web
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	//  (1) Viewer
	//  (2) Editor
	//  (3) Deletor
	r.GET("/view/:dt/:region", handler.ViewHandler) // view/20200101/강남구
	r.GET("/edit/:dt/:region", handler.EditHandler)
	r.POST("/delete/:dt/:region", handler.DeleteHandler)
	r.POST("/save/:dt/:region", handler.SaveHandler)

	r.Run(":8080")

	defer c.Stop()
}
