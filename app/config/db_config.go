package config

import (
	"context"
	"fmt"
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

	dns := cfg.FormatDSN()
	db, err := gorm.Open(gormMysqlDriver.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Fatal(context.Background(), fmt.Sprintf("Error initializing database on dns %s", dns), err)
		return nil, err
	}

	return db, nil
}
