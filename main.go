package main

import (
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/httpserver"
	"game-app/repository/migrator"
	"game-app/repository/mysql"
	"game-app/repository/mysql/mysqlaccesscontrol"
	"game-app/repository/mysql/mysqluser"
	"game-app/repository/redis/redismatching"
	"game-app/scheduler"
	"game-app/service/authorizationservice"
	"game-app/service/authservice"
	"game-app/service/backofficeuserservice"
	"game-app/service/matchingservice"
	"game-app/service/user"
	"game-app/validator/matchingvalidator"
	"game-app/validator/uservalidator"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
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
	cfg := config.Load()
	log.Printf("cfg : %+v", cfg)

	//TODO - add command for migration
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	// TODO - add struct and add this returned items as struct field
	authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingValidator, matchingSvc := setupServices(*cfg)

	go func() {
		server := httpserver.New(*cfg, authSvc, userSvc, backofficeUserSvc, authorizationSvc, userValidator, matchingSvc, matchingValidator)

		server.Serve()
	}()

	done := make(chan bool)

	go func() {
		sch := scheduler.New()
		sch.Start(done)
	}()

	terminate := make(chan os.Signal)
	signal.Notify(terminate)
	<-terminate
	log.Printf("received interrupt signal, shutting down gracefully...")
	done <- true
	time.Sleep(5 * time.Second)

}

func setupServices(cfg config.Config) (authservice.Service, user.Service, uservalidator.Validator, backofficeuserservice.Service, authorizationservice.Service, matchingvalidator.Validator, matchingservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userMysql := mysqluser.New(MysqlRepo)

	userSvc := user.New(authSvc, userMysql)

	aclMysql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMysql)

	backofficeUserSvc := backofficeuserservice.New()

	uV := uservalidator.New(userMysql)

	matchingV := matchingvalidator.New()
	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingConfig, matchingRepo)

	return authSvc, *userSvc, uV, backofficeUserSvc, authorizationSvc, matchingV, matchingSvc

}
