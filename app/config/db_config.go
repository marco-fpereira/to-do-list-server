package config

import (
	"context"
	"os"

	"github.com/marco-fpereira/to-do-list-server/config/logger"

	goSqlDriver "github.com/go-sql-driver/mysql"
	gormMysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect() (*gorm.DB, error) {
	m := make(map[string]string)
	m["charset"] = "utf8"

	cfg := goSqlDriver.Config{
		Addr:      os.Getenv("HOST") + ":" + os.Getenv("DBPORT"),
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		DBName:    os.Getenv("DBNAME"),
		Net:       "tcp",
		ParseTime: true,
		Params:    m,
	}

	db, err := gorm.Open(gormMysqlDriver.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		logger.Fatal(context.Background(), "Error initializing database", err)
		return nil, err
	}

	return db, nil
}
