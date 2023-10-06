package scheduler

import (
	"fmt"
	"log"
	"time"
)

type Scheduler struct {
}

func New() Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start(done <-chan bool) {
	for {
		select {
		case <-done:
			log.Printf("exiting...")
			return
		default:
			now := time.Now()
			fmt.Println("scheduler now", now)
			time.Sleep(time.Second)
		}
	}
}
