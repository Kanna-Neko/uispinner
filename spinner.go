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
	depth         int
	p             *Spinner
	child         []*Spinner
	childDoneNum  int
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
func (s *Spinner) AddSpinner(in []string, interval time.Duration) *Spinner {
	var new = NewSpinner(in, interval)
	new.belong = s.belong
	new.p = s
	new.depth = s.depth + 1
	s.child = append(s.child, new)
	s.Work()
	return new
}
func (s *Spinner) Work() {
	if !s.done {
		return
	}
	s.mtx.Lock()
	s.done = false
	s.mtx.Unlock()
	if s.depth > 0 && s.p.done {
		s.p.Work()
	}
}
func (s *Spinner) Bind(p *Process) *Spinner {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.belong = p
	return s
}
func (s *Spinner) Done() {
	if s.done {
		return
	}
	s.mtx.Lock()
	s.done = true
	s.mtx.Unlock()
	if s.depth != 0 {
		s.mtx.Lock()
		s.p.childDoneNum++
		s.mtx.Unlock()
		if s.p.childDoneNum == len(s.p.child) {
			s.p.Done()
		}
	}
	for _, v := range s.child {
		v.Done()
	}
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
	// if !s.done {
	// 	s.belong.RefreshInterval()
	// }
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
func (s *Spinner) String(front []bool) string {
	var res string
	if s.done {
		res = s.Complete + "\n"
	} else {
		res = s.prefix + s.SpinnerString[s.current] + s.suffix + "\n"
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
	}
	var pre string
	for _, v := range front {
		if v {
			pre += "│ "
		} else {
			pre += "  "
		}
	}
	for i, v := range s.child {
		if i == len(s.child)-1 {
			front = append(front, false)
			res += pre + "└─" + v.String(front)
			front = front[:len(front)-1]
		} else {
			front = append(front, true)
			res += pre + "├─" + v.String(front)
			front = front[:len(front)-1]
		}
	}
	return res
}
