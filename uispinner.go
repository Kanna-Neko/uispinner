package uispinner

import (
	"fmt"
	"io"
	"snipper/tool"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

type Progress struct {
	Spinners        []*Spinner
	lw              *uilive.Writer
	tdone           chan bool
	mtx             *sync.RWMutex
	refreshInterval time.Duration
}

func New() *Progress {
	return &Progress{
		Spinners:        make([]*Spinner, 0),
		lw:              uilive.New(),
		tdone:           make(chan bool),
		mtx:             &sync.RWMutex{},
		refreshInterval: 0,
	}
}

func (p *Progress) AddSpinner(stringSet []string, interval time.Duration) *Spinner {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	var res = NewSpinner(stringSet, interval).Bind(p)
	p.refreshInterval = time.Duration(tool.Gcd(int64(p.refreshInterval), int64(res.interval)))
	p.Spinners = append(p.Spinners, res)
	return res
}

func (p *Progress) SetRefreshInterval(interval time.Duration) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.refreshInterval = interval
}

func (p *Progress) Listen() {
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

func (p *Progress) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	for _, Spinner := range p.Spinners {
		fmt.Fprintln(p.lw, Spinner.String())
	}
	p.lw.Flush()
}

func (p *Progress) Start() {
	go p.Listen()
}

func (p *Progress) Stop() {
	p.tdone <- true
	<-p.tdone
}

func (p *Progress) Bypass() io.Writer {
	return p.lw.Bypass()
}
