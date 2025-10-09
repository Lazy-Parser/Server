// Package process
// To build this package, I borrowed some design ideas from those folks https://github.com/stretchr/testify
package process

import (
	"context"
	"time"
)

// 1. Maybe make *Process in Timer.
// 2. Make Do() function in Timer non-blocking + add context for cancellation

type Process struct {
	id      string // title
	startAt time.Time
	status  Status

	ctx    context.Context
	cancel context.CancelFunc

	p Processable
}

// initProcess - set up all internal things. Call in the initialization part
func (process *Process) initProcess(id string) {
	process.id = id
	process.startAt = time.Now()
	process.status = Waiting
	ctx, cancel := context.WithCancel(context.Background())
	process.ctx = ctx
	process.cancel = cancel
}

func (process *Process) GetID() string                 { return process.id }
func (process *Process) GetStartAt() time.Time         { return process.startAt }
func (process *Process) GetStatus() Status             { return process.status }
func (process *Process) GetCtx() context.Context       { return process.ctx }
func (process *Process) GetCancel() context.CancelFunc { return process.cancel }

func (process *Process) setStatus(status Status) { process.status = status }

// maybe add setters and decide to make them public / private
