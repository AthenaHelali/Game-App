package main

import (
	"context"
	"game-app/internal/adapter/redis"
	config2 "game-app/internal/config"
	"game-app/internal/delivery/httpserver"
	"game-app/internal/repository/mysql"
	"game-app/internal/repository/mysql/mysqlaccesscontrol"
	"game-app/internal/repository/mysql/mysqluser"
	"game-app/internal/repository/redis/redismatching"
	"game-app/internal/repository/redis/redispresence"
	"game-app/internal/scheduler"
	"game-app/internal/service/authorizationservice"
	"game-app/internal/service/authservice"
	"game-app/internal/service/backofficeuserservice"
	"game-app/internal/service/matchingservice"
	"game-app/internal/service/presenceservice"
	"game-app/internal/service/user"
	"game-app/internal/validator/matchingvalidator"
	"game-app/internal/validator/uservalidator"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
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
	cfg := config2.Load()
	log.Printf("cfg : %+v", cfg)

	//TODO - add command for migration
	//mgr := migrator.New(cfg.Mysql)
	//mgr.Up()

	// TODO - add struct and add this returned items as struct field
	authSvc, userSvc, userValidator, backofficeUserSvc, authorizationSvc, matchingSvc, matchingValidator, presenceSvc := setupServices(*cfg)

	server := httpserver.New(*cfg, authSvc, userSvc, backofficeUserSvc, authorizationSvc, userValidator,
		matchingSvc, matchingValidator, presenceSvc)
	go func() {
		server.Serve()
	}()

	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		sch := scheduler.New(cfg.Scheduler, matchingSvc)
		wg.Add(1)
		sch.Start(done, &wg)
	}()

	terminate := make(chan os.Signal)
	signal.Notify(terminate)
	<-terminate

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		log.Printf("httpserver shutdown error")
	}

	log.Printf("received interrupt signal, shutting down gracefully...")
	done <- true

	<-ctxWithTimeout.Done()

	wg.Wait()

}

func setupServices(cfg config2.Config) (
	authservice.Service, user.Service, uservalidator.Validator,
	backofficeuserservice.Service, authorizationservice.Service,
	matchingservice.Service, matchingvalidator.Validator,
	presenceservice.Service) {
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

	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(cfg.PresenceService, presenceRepo)

	return authSvc, *userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV, presenceSvc

}
