#GO-WORKER-POOL [![Hits](https://hits.sh/github.com/nawazish-github/go-worker-pool.svg)](https://hits.sh/github.com/nawazish-github/go-worker-pool/)

##Philosophy

Go-routines are cheap user space "threads". By cheap, I mean it is cheaper on the Golang runtime to create and destroy them; comparatively they take less memory overhead to be created. That is the reason, there are production system which work with tens of thousands of Go-routines at any given time.

However, "cheap" does not mean Go-routines can be used **freely**; they incurr cost albeit less. On internet scale architecture, these costs may add up to hamper systems. On such systems it is a good practice to bound the number of Go-routines to a certain emperically calculated number. Generally this number is calculated considering two factors:

1. Number of available CPU cores on the machine
2. Nature of the tasks; basically, the compute to I/O ratio.

##go-worker-pool

go-worker-pool creates a managed pool of Go-routines, called as workers,  which can be reused to schedule command-type (command/query pattern) tasks on it. In case if there are no available worker Go-routines to serve the request, the submission of the task would temporarily block unless there is a worker available in the pool again. 

The public APIs are available to create, submit tasks and shutdown the pool.

##Sample usage

```
import (
	"fmt"
	"time"

    "github.com/nawazish-github/go-worker-pool/pool"
)

type task struct {
	name string
}

func (t *task) Run() {
	fmt.Println(t.name)
	time.Sleep(time.Millisecond * 500)
}

func main() {
	p := pool.New(5)
	names := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj"}
	for i := 0; i < 10; i++ {
		p.Submit(&task{name: names[i]})
	}
    p.Shutdown()
	time.Sleep(time.Second * 5)
}
```