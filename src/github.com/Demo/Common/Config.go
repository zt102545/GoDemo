package Common

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	App appConfig `json:"app"`
	DataBase DatabaseConfig `json:"database"`
	Redis RedisConfig `json:"Redis"`
}

type appConfig struct {
	AppName string `json:"name"`
	AppMode string `json:"mode"`
	AppHost string `json:"host"`
	AppPort string `json:"port"`
}

type DatabaseConfig struct{
	Url string `json:"url"`
}

type RedisConfig struct{
	Addr string `json:"addr"`
	DB int `json:"db"`
	PassWord string `json:"password"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
