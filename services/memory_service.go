package services

import (
	"fmt"
	"github.com/umitdogan0/CustomCacheApp/configuration"
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
