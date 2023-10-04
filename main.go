package main

import (
	"fmt"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/mysqlaccesscontrol"
	"game-app/repository/mysql/mysqluser"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backofficeuserservice"
	"game-app/service/user"
	"game-app/validator/uservalidator"
	"os"
	"strconv"
	"time"
)

const (
	jwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func GetHTTPServerPort(fallback int) int {
	portStr := os.Getenv("GAMEAPP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fallback
	}

	return port
}
func main() {
	cfg2 := config.Load()
	fmt.Printf("cfg : %+v", cfg2)
	//merge cfg2 and cfg
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: GetHTTPServerPort(8088)},
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
			DBName:   "gameapp_db",
		},
	}

	//TODO - add command for migration
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc, backofficeUserSvc, authorizationSvc, userValidator)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, user.Service, uservalidator.Validator, backofficeuserservice.Service, authorizationservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userMysql := mysqluser.New(MysqlRepo)

	userSvc := user.New(authSvc, userMysql)

	aclMysql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMysql)

	backofficeUserSvc := backofficeuserservice.New()

	uV := uservalidator.New(userMysql)

	return authSvc, *userSvc, uV, backofficeUserSvc, authorizationSvc

}
