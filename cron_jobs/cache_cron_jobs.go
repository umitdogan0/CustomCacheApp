package cron_jobs

import (
	"github.com/umitdogan0/CustomCacheApp/services"
	"time"
)

func CheckCacheExpr() {
	for {
		services.CheckExpirationDate()
		time.Sleep(1 * time.Second)
	}
}
