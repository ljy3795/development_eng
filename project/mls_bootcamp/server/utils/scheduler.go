package utils

import (
	"github.com/robfig/cron"
)

// DailyCronAPI is scheduler for daily API call(utils.DailyAPICall)
func DailyCronAPI(cronInput string) {
	c := cron.New()
	c.AddFunc(cronInput, func() {
		DailyAPICall()
	})
	c.Start()

	// defer c.Stop()
}
