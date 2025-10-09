package process

import (
	"context"
	"log"
	"time"
)

type Timer struct {
	Process

	timeNow   time.Time
	isRunning bool
}

func (timer *Timer) SetupProcess() {
	log.Printf("SETUP FUNCTION BEFORE '%s' PROCESS\n", timer.id)
}

func NewTimer(title string) *Timer {
	timer := &Timer{
		timeNow:   time.Now(),
		isRunning: true,
	}
	timer.initProcess(title)
	return timer
}

func (timer *Timer) Do(ctx context.Context) {
	for {
		if timer.isRunning {
			time.Sleep(time.Second)
			timer.timeNow = timer.timeNow.Add(time.Second)
		}
	}
}

func (timer *Timer) StopProcess(cancel context.CancelFunc) {
	timer.isRunning = false
	timer.setStatus(Completed)
}

func (timer *Timer) GetTime() string {
	return timer.timeNow.Format(time.TimeOnly)
}
