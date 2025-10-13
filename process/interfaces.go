package process

import (
	"context"
	"time"
)

// Processable has a Do function, which is a main function for any task.
type Processable interface {
	Do(ctx context.Context)

	GetID() string
	GetStartAt() time.Time
	GetStatus() Status
	GetCtx() context.Context
	GetCancel() context.CancelFunc

	setStatus(status Status)
}

// SetupProcess has a SetupProcess method, which will run before the Do function
type SetupProcess interface {
	SetupProcess()
}

// StopProcess has a StopProcess method, calling which the Do will stop
type StopProcess interface {
	StopProcess(cancel context.CancelFunc)
}

// StatsProcess TODO: implement
type StatsProcess interface {
}
