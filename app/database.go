package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "sql6507871:N8ArIFBNIn(sql6.freemysqlhosting.net:3306)/sql6507871")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
