package sdcp

import (
	"errors"
	"sync"

	"github.com/loopholelabs/logging/types"
)

var (
	ErrAlreadyRegistered = errors.New("machine already registered")
	ErrRegisterFailed    = errors.New("failed to register machine")
)

type SDCP struct {
	logger types.Logger

	machinesMu sync.RWMutex
	machines   map[string]*Machine
}

func New(logger types.Logger) *SDCP {
	return &SDCP{
		logger:   logger.SubLogger("sdcp"),
		machines: make(map[string]*Machine),
	}
}

func (s *SDCP) Register(machineID string, machineIP string) error {
	s.machinesMu.Lock()
	if _, ok := s.machines[machineID]; ok {
		s.machinesMu.Unlock()
		return ErrAlreadyRegistered
	}
	m, err := newMachine(machineID, machineIP, s.logger)
	if err != nil {
		s.machinesMu.Unlock()
		return errors.Join(ErrRegisterFailed, err)
	}
	s.machines[machineID] = m
	s.machinesMu.Unlock()
	return nil
}

func (s *SDCP) Unregister(machineID string) {
	s.machinesMu.Lock()
	if m, ok := s.machines[machineID]; ok {
		m.stop()
		delete(s.machines, machineID)
	}
	s.machinesMu.Unlock()
}

func (s *SDCP) GetMachine(machineID string) (*Machine, bool) {
	s.machinesMu.RLock()
	m, ok := s.machines[machineID]
	s.machinesMu.RUnlock()
	return m, ok
}

func (s *SDCP) Close() {
	s.machinesMu.Lock()
	for _, m := range s.machines {
		m.stop()
	}
	clear(s.machines)
	s.machinesMu.Unlock()
}
