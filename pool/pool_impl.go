package pool

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/nawazish-github/go-worker-pool/runnable"
)

const poolShutdown = 1

type poolImpl struct {
	runnableCh chan runnable.Runnable
	numWrkrs   int32
	freeWrkrs  int32
	exit       chan struct{}
	isShutdown int32
}

// New creates a new instance of the pool and returns it.
// Calling New multiple times or concurrently would return
// New instance per call.
func New(numWrkrs int32) Pool {
	p := &poolImpl{
		runnableCh: make(chan runnable.Runnable),
		numWrkrs:   numWrkrs,
		freeWrkrs:  numWrkrs,
		exit:       make(chan struct{}),
		isShutdown: 0,
	}

	for i := 0; i < int(numWrkrs); i++ {
		go p.workerStart(i)
	}
	return p
}
func (p *poolImpl) Submit(r runnable.Runnable) error {
	if atomic.LoadInt32(&p.isShutdown) != poolShutdown {
		p.runnableCh <- r
		return nil
	}
	return errors.New("CANNOT SUBMIT TASK ON A POOL WHICH IS SHUT")

}
func (p *poolImpl) Shutdown() {
	if shutdown := atomic.CompareAndSwapInt32(&p.isShutdown, 0, 1); shutdown {
		n := atomic.LoadInt32(&p.numWrkrs)
		for i := 0; i < int(n); i++ {
			p.exit <- struct{}{}
		}
		fmt.Println("pool shut down complete. free workers: ", atomic.LoadInt32(&p.freeWrkrs))
	} else {
		fmt.Println("Pool is already shut")
	}

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
