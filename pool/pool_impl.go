package pool

import (
	"fmt"
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
		go p.workerStart(i)
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
	fmt.Println("shut down complete. free workers: ", atomic.LoadInt32(&p.freeWrkrs))
}

func (p *poolImpl) workerStart(id int) {
	for {
		select {
		case <-p.exit:
			atomic.AddInt32(&p.freeWrkrs, -1)
			return
		case w := <-p.runnableCh:
			atomic.AddInt32(&p.freeWrkrs, -1)
			fmt.Printf("go-routine id: %d ", id)
			w.Run()
			atomic.AddInt32(&p.freeWrkrs, 1)
		}
	}
}
