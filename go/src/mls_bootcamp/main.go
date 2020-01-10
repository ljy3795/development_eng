package main

import (
	"io"
	"log"
	"net/http"
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
	strd_dt := time.Now() // cron 시작일

	sta_i := 1    // 시작페이지
	end_i := 1000 // 끝페이지

	_, _, rows := mylibs.GetFromAPI(sta_i, end_i, strd_dt)

	// Load to MariaDB
	// 1) db Conn 생성
	db := mylibs.DbConn()
	defer db.Close()

	// 2) DB에 적재
	ins_rows := mylibs.ReplaceQuery(db, rows)
	myLogger.Printf("%d row inserted.\n", ins_rows)
	myLogger.Println(strd_dt.Format("20060102"), " is done")
}

func main() {
	myLogger = log.New(os.Stdout, "Info: ", log.LstdFlags)

	// 0) Backfill
	strd_dt := time.Now().AddDate(0, 0, -1)
	backfill_cnt := 1

	for i := 0; i < backfill_cnt; i++ {
		//   Args.
		sta_i := 1    // 시작 페이지
		end_i := 1000 // 끝 페이지

		// API Call & GET result
		// -> strd_dt 일자
		_, _, rows := mylibs.GetFromAPI(sta_i, end_i, strd_dt)

		// Load to MariaDB
		// 1) db Conn 생성
		db := mylibs.DbConn()
		defer db.Close()

		ins_rows := mylibs.ReplaceQuery(db, rows)
		myLogger.Printf("%d row inserted.\n", ins_rows)
		myLogger.Println(strd_dt.Format("20060102"), " is done")
		strd_dt = strd_dt.AddDate(0, 0, -1)
		db.Close()
	}

	// 1) Scheduler for daily API Call (everyday on 9am)
	c := cron.New()
	c.AddFunc("0 8 * * *", func() {
		dailyAPICall()
	})
	go c.Start()
	// sig := make(chan os.Signal)
	// signal.Notify(sig, os.Interrupt, os.Kill)
	// <-sig
	c.Stop()

	// 2) Web Handler
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// web
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	//  (1) Viewer
	//  (2) Editor
	//  (3) Deletor
	r.GET("/view/:dt/:region", handler.ViewHandler) // view/20200101/강남구
	r.POST("/edit/:dt/:region", handler.EditHandler)
	r.POST("/save/:dt/:region", handler.SaveHandler)

	r.Run(":8080")
}
