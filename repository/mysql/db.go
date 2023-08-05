package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	db *sql.DB
}

func New() *MysqlDB {
	db, err := sql.Open("mysql", "gameapp:gameappt0lk2o20@(localhost:3308)/gameapp_db")

	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDB{db}
}
