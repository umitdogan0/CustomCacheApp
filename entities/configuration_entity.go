package entities

type ConfigEntity struct {
	Server struct {
		MaximumMemory                int    `yaml:"maximum_memory""`
		Port                         string `yaml:"port"`
		AutomaticCleaning            bool   `yaml:"automatic_cleaning"`
		MinAutomaticCleaningPriority int    `yaml:"min_automatic_cleaning_priority"`
		MaxAutomaticCleaningPriority int    `yaml:"max_automatic_cleaning_priority"`
	} `yaml:"server"`
}
