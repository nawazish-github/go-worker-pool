package main

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
