package helpers

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// Config for configuration type
type Config struct {
	Name        string
	Version     string
	Debug       bool
	Port        string
	SignKey     string
	SQLPath     string
	StaticPath  string
	BaseAuth    BaseAuth
	MysqlConfig MysqlConfig
}

// BaseAuth realm base auth
type BaseAuth struct {
	User     string
	Password string
}

// MysqlConfig for mysql configuration type
type MysqlConfig struct {
	Host       string
	User       string
	Password   string
	Port       string
	Database   string
	MaxRetries int
}

// ReadConfig read conf file
func ReadConfig(configFilePath string) Config {

	var returnConfig Config

	// Decode TOML file data to data structure
	if meta, err := toml.DecodeFile(configFilePath, &returnConfig); err != nil {
		log.WithFields(log.Fields{
			"ConfigFile": configFilePath,
			"Error":      err,
		}).Error("Error decoding TOML file")
	} else {
		printConfiguration(configFilePath, returnConfig, meta)
	}

	return returnConfig
}

// printConfiguration
func printConfiguration(filePath string, conf Config, meta toml.MetaData) {
	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		log.WithFields(log.Fields{
			"ConfigFilePath": filePath,
			"Undecoded":      meta.Undecoded(),
		}).Warn("Unable to decode values in config")
	}
}
