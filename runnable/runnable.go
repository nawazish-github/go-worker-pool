package runnable

// Runnable abstracts the task that can be scheduled
// on the Worker Pool. The nature of the runnable task
// is that of command type, in it, it does not return
// any value (neither does it take any argument)
type Runnable interface {
	Run()
}