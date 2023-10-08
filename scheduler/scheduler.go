package scheduler

import (
	"game-app/param"
	"game-app/service/matchingservice"
	"github.com/go-co-op/gocron"
	"log"
	"sync"
	"time"
)

type Config struct {
	MatchWaitingUsersIntervalInSeconds int `koanf:"match_waiting_users_interval_in_seconds"`
}
type Scheduler struct {
	config   Config
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(config Config, matchSvc matchingservice.Service) Scheduler {
	return Scheduler{config: config, sch: gocron.NewScheduler(time.UTC), matchSvc: matchSvc}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	s.sch.Every(s.config.MatchWaitingUsersIntervalInSeconds).Second().Do(s.MatchWaitingUsers)

	s.sch.StartAsync()

	<-done
	log.Printf("stop scheduler...")
	s.sch.Stop()
}
func (s Scheduler) MatchWaitingUsers() {
	s.matchSvc.MatchWaitingUsers(param.MatchWaitingUsersRequest{})

}
