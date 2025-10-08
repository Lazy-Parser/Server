package process

type Status int

const (
	Running Status = iota
	Waiting        // Set this state at initialization, but before running. Meaning, that is waiting to start
	Failed
	Completed
)

var stateName = map[Status]string{
	Running:   "running",
	Waiting:   "waiting",
	Failed:    "failed",
	Completed: "completed",
}

func (s Status) String() string {
	return stateName[s]
}
