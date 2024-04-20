package entities

import "time"

type CacheEntity struct {
	Data       map[string]interface{}
	Expiration int
	Priority   int
	CreateDate time.Time
}
