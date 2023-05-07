package util

import (
	"encoding/json"
	"io"
	"os"
)

type DBCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerCfg struct {
	Env  string `json:"env"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Cfg struct {
	Server ServerCfg `json:"server"`
	DB     DBCfg     `json:"db"`
}

func LoadCfg(filename string) (*Cfg, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := &Cfg{}
	json.Unmarshal([]byte(byteData), cfg)

	return cfg, nil
}
