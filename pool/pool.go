package pool

import "github.com/nawazish-github/go-worker-pool/runnable"

// Pool is  the worker pool that contains a managed
// pool of worker go-routines. The pool can be created
// by choosing right amount of go-routines using the 
// pool.New() function.
//
// Then runnable.Runnable tasks can be scheduled on the pool's 
// available worker using the Submit API.
//
// The worker pool can be shutdown gracefully using the Shutdown 
// method on this interface. This would allow the existing tasks 
// to complete before the pool is shutdown.
//
// Once shutdown has been called on the pool no tasks can be 
// further submitted on it. Any such submission would be result
// in error
type Pool interface {
	Submit(r runnable.Runnable) error
	Shutdown()
}