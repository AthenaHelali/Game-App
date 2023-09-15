package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DNName   string `koanf:"DNName"`
}

type MysqlDB struct {
	config Config
	db     *sql.DB
}

func New(cfg Config) *MysqlDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DNName))

	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDB{config: cfg, db: db}
}
