package process

import (
	"errors"
	"log"
)

type Manager struct {
	processes map[string]Processable
}

func NewProcessManager() *Manager {
	return &Manager{
		processes: make(map[string]Processable),
	}
}

func (processManager *Manager) Get(id string) (Processable, bool) {
	p, ok := processManager.processes[id]
	return p, ok
}

func (processManager *Manager) GetList() map[string]Processable {
	return processManager.processes
}

func (processManager *Manager) Append(process Processable) error {
	if _, exists := processManager.Get(process.GetID()); exists {
		return errors.New("process already exists")
	}

	if setupProcess, ok := process.(SetupProcess); ok {
		if err := setupProcess.SetupProcess(process.GetCtx()); err != nil {
			panic(err)
		}
	}

	// start process
	process.setStatus(Running)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover from panic: %v", r)
		}
	}()
	go process.Do(process.GetCtx())

	processManager.processes[process.GetID()] = process
	return nil
}

func (processManager *Manager) Stop(id string) bool {
	if p, exists := processManager.Get(id); exists {
		if cancellable, ok := p.(StopProcess); ok {
			cancellable.StopProcess(p.GetCancel())
			return true
		}
	}
	return false
}
