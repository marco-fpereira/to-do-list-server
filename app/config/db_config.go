package config

import (
	"os"

	goSqlDriver "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	gormMysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect() (*gorm.DB, error) {
	cfg := goSqlDriver.Config{
		Addr:   os.Getenv("HOST") + ":" + os.Getenv("DBPORT"),
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
		Net:    "tcp",
	}

	db, err := gorm.Open(gormMysqlDriver.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
