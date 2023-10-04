package mysqluser

import "game-app/repository/mysql"

type DB struct {
	conn *mysql.MysqlDB
}

func New(conn *mysql.MysqlDB) *DB {
	return &DB{conn: conn}
}
