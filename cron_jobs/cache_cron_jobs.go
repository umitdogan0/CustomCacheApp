package cron_jobs

import (
	"custom_cache_app/services"
	"time"
)

func CheckCacheExpr() {
	for {
		services.CheckExpirationDate()
		time.Sleep(1 * time.Second)
	}
}
