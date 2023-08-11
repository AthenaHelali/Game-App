package main

import (
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/mysql"
	"game-app/service/authservice"
	"game-app/service/user"
	"time"
)

const (
	jwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {

	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8088},
		Auth: authservice.Config{
			SignKey:               jwtSignKey,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3308,
			Host:     "localhost",
			DNName:   "gameapp_db",
		},
	}
	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, user.Service) {
	authSvc := authservice.New(cfg.Auth)
	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := user.New(authSvc, MysqlRepo)

	return authSvc, *userSvc

}
