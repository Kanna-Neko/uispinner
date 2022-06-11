package uispinner

import (
	"sync"
	"time"
)

type Spinner struct {
	prefix        string
	suffix        string
	Complete      string
	SpinnerString []string
	done          bool
	interval      time.Duration
	current       int
	currentTime   time.Duration
	belong        *Process
	mtx           *sync.RWMutex
}

func NewSpinner(in []string, interval time.Duration) *Spinner {
	return &Spinner{
		SpinnerString: in,
		done:          false,
		current:       0,
		interval:      interval,
		mtx:           &sync.RWMutex{},
	}
}
func (s *Spinner) Bind(p *Process) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.belong = p
	return s
}
func (s *Spinner) Done() *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.done = true
	s.belong.RefreshInterval()
	return s
}
func (s *Spinner) SetCharSet(in []string) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.SpinnerString = in
	s.current = 0
	return s
}
func (s *Spinner) SetInterval(interval time.Duration) *Spinner {
	s.mtx.Lock()
	s.interval = interval
	s.mtx.Unlock()
	if !s.done {
		s.belong.RefreshInterval()
	}
	return s
}
func (s *Spinner) SetComplete(in string) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.Complete = in
	return s
}
func (s *Spinner) Reverse() *Spinner {
	i := 0
	j := len(s.SpinnerString) - 1
	for i < j {
		s.SpinnerString[i], s.SpinnerString[j] = s.SpinnerString[j], s.SpinnerString[i]
		i++
		j--
	}
	return s
}
func (s *Spinner) SetPrefix(in string) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.prefix = in
	return s
}
func (s *Spinner) SetSuffix(in string) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.suffix = in
	return s
}
func (s *Spinner) String() string {
	if s.done {
		return s.Complete
	} else {
		defer func() {
			s.mtx.Lock()
			defer s.mtx.Unlock()
			s.currentTime += s.belong.refreshInterval
			if s.currentTime >= s.interval {
				s.current++
				s.current %= len(s.SpinnerString)
				s.currentTime %= s.interval
			}
		}()
		return s.prefix + s.SpinnerString[s.current] + s.suffix
	}
}
