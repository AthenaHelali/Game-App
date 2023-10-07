package scheduler

import (
	"game-app/param"
	"game-app/service/matchingservice"
	"github.com/go-co-op/gocron"
	"log"
	"sync"
	"time"
)

type Scheduler struct {
	sch      *gocron.Scheduler
	matchSvc matchingservice.Service
}

func New(matchSvc matchingservice.Service) Scheduler {
	return Scheduler{sch: gocron.NewScheduler(time.UTC), matchSvc: matchSvc}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	s.sch.Every(5).Second().Do(s.MatchWaitingUsers)

	s.sch.StartAsync()

	<-done
	log.Printf("stop scheduler...")
	s.sch.Stop()
}
func (s Scheduler) MatchWaitingUsers() {
	s.matchSvc.MatchWaitingUsers(param.MatchWaitingUsersRequest{})

}
