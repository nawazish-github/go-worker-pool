package pool

import (
	"sync"
	"sync/atomic"

	"github.com/nawazish-github/go-worker-pool/runnable"
)

type poolImpl struct {
	runnableCh chan runnable.Runnable
	numWrkrs   int32
	freeWrkrs  int32
	exit       chan struct{}
}

func New(numWrkrs int32) Pool {
	p := &poolImpl{
		runnableCh: make(chan runnable.Runnable),
		numWrkrs:   numWrkrs,
		freeWrkrs:  numWrkrs,
		exit:       make(chan struct{}),
	}

	for i := 0; i < int(numWrkrs); i++ {
		go p.workerStart()
	}
	return p
}
func (p *poolImpl) Submit(r runnable.Runnable) {
	p.runnableCh <- r
}
func (p *poolImpl) Shutdown() {
	n := atomic.LoadInt32(&p.numWrkrs)
	for i := 0; i < int(n); i++ {
		p.exit <- struct{}{}
	}
}

func (p *poolImpl) workerStart() {
	for {
		select {
		case <-p.exit:
			return
		case w := <-p.runnableCh:
			atomic.AddInt32(&p.freeWrkrs, -1)
			w.Run()
			atomic.AddInt32(&p.freeWrkrs, 1)
		}
	}
}
