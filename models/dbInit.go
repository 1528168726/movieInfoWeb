package models

import (
	"database/sql"
	"databaseHW/config"
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit(dbInfo *config.DbInfo) error {
	connStr := "host=" + dbInfo.Host + " port=" + dbInfo.Port + " user=" + dbInfo.User + " password=" + dbInfo.Password +
		" dbname=" + dbInfo.DBName + " sslmode=disable"
	rawdb, err := sql.Open("opengauss", connStr)
	if err != nil {
		return err
	}
	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: rawdb,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
