package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppParams Params           `json:"app_params"`
	DBParams  DBPostgresParams `json:"db_params"`
}

type Params struct {
	ServerURL    string `json:"server_url"`
	PortRun      string `json:"port"`
	WriteTimeout int64  `json:"write_timeout"`
	ReadTimeout  int64  `json:"read_timeout"`
}

type DBPostgresParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	DataBase string `json:"database"`
	Schema   string `json:"schema"`
}

func InitConfig() (*Config, error) {
	fmt.Println("Starting reading settings file")
	configFile, err := os.Open("./configs.json")
	if err != nil {
		return nil, fmt.Errorf("couldn't open config file: %v", err)
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
		}
	}(configFile)

	fmt.Println("Starting decoding settings file")

	var config Config
	if err = json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, fmt.Errorf("couldn't decode settings json file: %v", err)
	}

	return &config, nil
}
