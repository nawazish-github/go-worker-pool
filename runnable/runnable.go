package runnable

// Runnable abstracts the task that can be scheduled
// on the Worker Pool. The nature of the runnable task
// is that of command type, in it, it does not return
// any value (neither does it take any argument)
//
// Client should encapsulate their task complying to the
// signature of the Run() method inorder to submit it 
// for execution onto the worker pool.
type Runnable interface {
	Run()
}