package process

import "time"

// Processable has a Do function, which is a main function for any task.
type Processable interface {
	Do()
	GetID() string
	GetStartAt() time.Time
	GetStatus() Status

	setStatus(status Status)
}

// SetupProcess has a SetupProcess method, which will run before the Do function
type SetupProcess interface {
	SetupProcess()
}

// StopProcess has a StopProcess method, calling which the Do will stop
type StopProcess interface {
	StopProcess()
}

// StatsProcess TODO: implement
type StatsProcess interface {
}
