package pool

import "github.com/nawazish-github/go-worker-pool/runnable"

type Pool interface {
	Submit(r runnable.Runnable)
	Shutdown()
}