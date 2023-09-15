package config

import (
	"game-app/repository/mysql"
	"game-app/service/authservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	HTTPServer HTTPServer         `koanf:"http_server"`
	Auth       authservice.Config `koanf:"auth"`
	Mysql      mysql.Config       `koanf:"mysql"`
}
