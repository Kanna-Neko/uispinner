package uispinner

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

// responsible for docking with io and fresh io
type Process struct {
	Spinners        []*Spinner
	lw              *uilive.Writer
	tdone           chan bool
	mtx             *sync.RWMutex
	refreshInterval time.Duration
}

// return a new process
func New() *Process {
	return &Process{
		Spinners:        make([]*Spinner, 0),
		lw:              uilive.New(),
		tdone:           make(chan bool),
		mtx:             &sync.RWMutex{},
		refreshInterval: 50 * time.Millisecond,
	}
}

// add a spinner to process manager
func (p *Process) AddSpinner(stringSet []string, interval time.Duration) *Spinner {
	p.mtx.Lock()
	var res = newSpinner(stringSet, interval).bind(p)
	p.Spinners = append(p.Spinners, res)
	p.mtx.Unlock()
	// p.RefreshInterval()
	return res
}


// Process run in the background
func (p *Process) listen() {
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

// fresh io and print spinner info
func (p *Process) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	for _, Spinner := range p.Spinners {
		fmt.Fprint(p.lw, Spinner.String([]bool{}))
	}
	p.lw.Flush()
}

// process start work
func (p *Process) Start() {
	go p.listen()
}

// process stop work
func (p *Process) Stop() {
	p.tdone <- true
	<-p.tdone
}

// return a io.Writer
// Bypass creates an io.Writer which allows non-buffered output to be written to the underlying output
func (p *Process) Bypass() io.Writer {
	return p.lw.Bypass()
}
