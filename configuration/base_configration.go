package configuration

import (
	"errors"
	"github.com/umitdogan0/CustomCacheApp/entities"
	"gopkg.in/yaml.v3"
	"os"
)

var configuration entities.ConfigEntity

func SetInitialConfiguration() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return errors.New("config.yaml is not found!")
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&configuration)
	if err != nil {
		return errors.New("Error in decoding the config.yaml file!")
	}
	if configuration.Server.AutomaticCleaning && configuration.Server.MinAutomaticCleaningPriority > configuration.Server.MaxAutomaticCleaningPriority {
		return errors.New("MinAutomaticCleaningPriority can not be greater than MaxAutomaticCleaningPriority")

	}
	return nil
}
func GetConfiguration() entities.ConfigEntity {
	return configuration
}
