// Package worker provides a worker
// pool manager for concurrent processing
// of tasks.
// VERSION:       1.0.0
// CONTRIBUTORS:  Ringo Hoffmann
package worker

// Worker is the struct managing the task
// queue and results collector.
type Worker struct {
	tasks           chan *Task
	results         chan *Result
	queueLength     int
	ResultHandler   func(*Result)
	FinishedHandler func()
}

// NewWorker creates a new instance of worker.
// bufferSize will be the size of the tas queue and
// result channel. Defaultly, this value is set to 1000
// if no argument was passed.
func NewWorker(bufferSize ...int) *Worker {
	bSize := 1000
	if len(bufferSize) > 0 {
		bSize = bufferSize[0]
	}
	return &Worker{
		tasks:   make(chan *Task, bSize),
		results: make(chan *Result, bSize),
	}
}

// SetResultHandler sets the ResultHandler of the worker,
// which will be called each time a worker finished and
// returned a Result instance containing the value interface
// and an error instance, if the function failed.
func (w *Worker) SetResultHandler(resultHandler func(*Result)) {
	w.ResultHandler = resultHandler
}

// SetFinishedHandler sets the FinishedHandler of the worker,
// which will be called after all tasks in the queue are
// finished.
func (w *Worker) SetFinishedHandler(finishedHandler func()) {
	w.FinishedHandler = finishedHandler
}

// Enqueue enqueues a task to the task queue.
func (w *Worker) Enqueue(task *Task) {
	w.tasks <- task
	w.queueLength++
}

// Start starts the processing of the queue.
// workersAmmount defines how much workers will
// process on th tasks queue simultaneously.
// The function returns, when all tasks in the
// queue are finished.
func (w *Worker) Start(workersAmmount int) {
	close(w.tasks)

	for i := 0; i < workersAmmount; i++ {
		go func(id int) {
			for t := range w.tasks {
				w.results <- t.F(t.Args)
			}
		}(i)
	}

	if w.ResultHandler == nil {
		w.ResultHandler = func(*Result) {}
	}

	for r := 0; r < w.queueLength; r++ {
		w.ResultHandler(<-w.results)
	}

	close(w.results)

	if w.FinishedHandler != nil {
		w.FinishedHandler()
	}
}
