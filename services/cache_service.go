package services

import (
	"errors"
	"fmt"
	"github.com/umitdogan0/CustomCacheApp/configuration"
	"github.com/umitdogan0/CustomCacheApp/entities"
	"time"
	"unsafe"
)

var caches *[]entities.CacheEntity = &[]entities.CacheEntity{}

func MakeCache(entity *entities.CacheEntity, key string) error {

	GetCollectionAlloc()
	checked := CheckMemoryUsage()
	if a := checkSameKey(key); a > -1 {
		(*caches)[a] = *entity
		return nil
	}
	fmt.Println(checked)
	if checked {

		*caches = append(*caches, *entity)
		return nil
	}
	config := configuration.GetConfiguration()
	if config.Server.AutomaticCleaning {
		err := cleanCache()
		if err != nil {
			return err
		}
		*caches = append(*caches, *entity)
	} else {
		return errors.New("Not enough memory")
	}
	return nil
}

func cleanCache() error {
	success := false
	var config = configuration.GetConfiguration().Server
	for i, cache := range *caches {
		if cache.Priority >= config.MinAutomaticCleaningPriority && cache.Priority <= config.MaxAutomaticCleaningPriority {
			*caches = append((*caches)[:i], (*caches)[i+1:]...)
			if CheckMemoryUsage() == true {
				success = true
				break
			}
		}
	}
	if !success {
		return errors.New("Not enough memory")
	}
	return nil
}

func checkSameKey(key string) int {
	for i, cache := range *caches {
		if _, ok := cache.Data[key]; ok {
			return i
		}
	}
	return -1
}

func CheckExpirationDate() {
	for i, cache := range *caches {
		if cache.Expiration == 0 {
			continue
		}
		if cache.CreateDate.Add(time.Duration(cache.Expiration) * time.Second).Before(time.Now()) {
			*caches = append((*caches)[:i], (*caches)[i+1:]...)
		}
	}
}

func GetCollectionAlloc() int {
	elementSize := unsafe.Sizeof(entities.CacheEntity{})
	alloc := len(*caches) * int(elementSize) / 1024 / 1024
	fmt.Println(len(*caches) * int(elementSize))
	return alloc
}

func GetCache(key string) (entities.CacheEntity, error) {
	for _, cache := range *caches {
		if _, ok := cache.Data[key]; ok {
			return cache, nil
		}
	}
	return entities.CacheEntity{}, errors.New("Cache not found")
}

func GetCaches() []entities.CacheEntity {
	return *caches
}
