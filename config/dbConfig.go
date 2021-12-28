package config

import (
	"encoding/json"
	"io"
	"os"
)

type DbInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDbInfo(fileName string) (DbInfo, error) {
	dbInfo := DbInfo{}
	fp, err := os.Open(fileName)
	if err != nil {
		return DbInfo{}, err
	}
	info, err := io.ReadAll(fp)
	if err != nil {
		return DbInfo{}, err
	}
	err = json.Unmarshal(info, &dbInfo)
	if err != nil {
		return DbInfo{}, err
	}
	return dbInfo, nil
}
