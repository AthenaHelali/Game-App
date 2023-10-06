package config

import (
	"game-app/adapter/redis"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/matchingservice"
)

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	HTTPServer     HTTPServer             `koanf:"http_server"`
	Auth           authservice.Config     `koanf:"auth"`
	Mysql          mysql.Config           `koanf:"mysql"`
	MatchingConfig matchingservice.Config `koanf:"matching_service"`
	Redis          redis.Config           `koanf:"redis"`
}
