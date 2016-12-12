package common

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

var AMCVersion string
var AMCBuild string
var AMCEdition string
var AMCEnv string

func AMCIsProd() bool {
	return AMCEnv == "prod"
}

func AMCIsEnterprise() bool {
	return AMCEdition == "enterprise"
}

type Config struct {
	AMC struct {
		UpdateInterval int    `toml:"update_interval"`
		CertFile       string `toml:"certfile"`
		KeyFile        string `toml:"keyfile"`
		StaticPath     string `toml:"static_dir"`

		Bind     string `toml:"bind"`
		LogLevel string `toml:"loglevel"`
		ErrorLog string `toml:"errorlog"`
		Chdir    string `toml:"chdir"`
		Timeout  int    `toml:"timeout"`
	}
}

func InitConfig(configFile, configDir string, config *Config) {
	// to print everything out regarding reading the config in app init
	log.SetLevel(log.DebugLevel)

	log.Info("Reading config file...")
	blob, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.Decode(string(blob), &config); err != nil {
		log.Fatal(err)
	}

	if config.AMC.Chdir != "" {
		if err := os.Chdir(config.AMC.Chdir); err != nil {
			log.Fatal("Error while trying to chdir to the specified directory in "+configFile+":", err)
		}
	}

	// setLogFile(config.AMC.ErrorLog)
	setLogLevel(config.AMC.LogLevel)
}

func setLogFile(filepath string) {
	out, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(out)
}

func setLogLevel(level string) {
	level = strings.ToLower(level)
	log.SetLevel(log.InfoLevel)
	switch level {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	}
}
