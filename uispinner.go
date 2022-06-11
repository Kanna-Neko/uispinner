package uispinner

import (
	"fmt"
	"io"
	"sync"
	"time"
	"uisnipper/tool"

	"github.com/gosuri/uilive"
)

type Process struct {
	Spinners        []*Spinner
	lw              *uilive.Writer
	tdone           chan bool
	mtx             *sync.RWMutex
	refreshInterval time.Duration
}

func New() *Process {
	return &Process{
		Spinners:        make([]*Spinner, 0),
		lw:              uilive.New(),
		tdone:           make(chan bool),
		mtx:             &sync.RWMutex{},
		refreshInterval: 0,
	}
}

func (p *Process) AddSpinner(stringSet []string, interval time.Duration) *Spinner {
	p.mtx.Lock()
	var res = NewSpinner(stringSet, interval).Bind(p)
	p.Spinners = append(p.Spinners, res)
	p.mtx.Unlock()
	p.RefreshInterval()
	return res
}

func (p *Process) RefreshInterval() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	var interval int64 = int64(time.Second)
	for i := 0; i < len(p.Spinners); i++ {
		if p.Spinners[i].done {
			continue
		}
		interval = tool.Gcd(interval, int64(p.Spinners[i].interval))
	}
	p.refreshInterval = time.Duration(interval)
}

func (p *Process) Listen() {
	for {
		p.mtx.Lock()
		interval := p.refreshInterval
		p.mtx.Unlock()

		select {
		case <-time.After(interval):
			p.print()
		case <-p.tdone:
			p.print()
			close(p.tdone)
			return
		}
	}
}

func (p *Process) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	for _, Spinner := range p.Spinners {
		fmt.Fprintln(p.lw, Spinner.String())
	}
	p.lw.Flush()
}

func (p *Process) Start() {
	go p.Listen()
}

func (p *Process) Stop() {
	p.tdone <- true
	<-p.tdone
}

func (p *Process) Bypass() io.Writer {
	return p.lw.Bypass()
}
