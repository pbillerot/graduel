/*
Configuration package is used to read the configuration file
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pbillerot/graduel/types"
)

var err error
var config types.Configuration
var about types.About

//ReadConfig will read the configuration json file to read the parameters
//which will be passed in the config file
func ReadConfig() (types.Configuration, error) {
	if config.ServerPort == "" {
		configFile, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Print("Unable to read config file, switching to flag mode")
			return types.Configuration{}, err
		}
		//log.Print(configFile)
		err = json.Unmarshal(configFile, &config)
		if err != nil {
			log.Print("Invalid JSON, expecting port from command line flag")
			return types.Configuration{}, err
		}
	}
	return config, nil
}

// ReadAbout c'est ça
func ReadAbout() (types.About, error) {
	if about.Application == "" {
		aboutFile, err := ioutil.ReadFile("about.json")
		if err != nil {
			log.Print("Unable to read config file, switching to flag mode")
			return types.About{}, err
		}
		//log.Print(configFile)
		err = json.Unmarshal(aboutFile, &about)
		if err != nil {
			log.Print("Invalid JSON, expecting port from command line flag")
			return types.About{}, err
		}
	}
	return about, nil
}
