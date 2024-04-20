package services

import (
	"custom_cache_app/configuration"
	"fmt"
)

func CheckMemoryUsage() bool {
	alloc := GetCollectionAlloc()
	config := configuration.GetConfiguration()
	fmt.Println(config.Server.MaximumMemory)
	if alloc > config.Server.MaximumMemory {
		return false
	}
	return true
}
