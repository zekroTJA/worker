package main

import (
	"fmt"
	"time"

	"github.com/zekroTJA/worker"
)

func main() {
	w := worker.NewWorker()

	for j := 0; j < 50; j++ {
		w.Enqueue(worker.NewTask(func(args []interface{}) *worker.Result {
			time.Sleep(1 * time.Second)
			return worker.NewResult(args[0], nil)
		}, j))
	}

	w.ResultHandler = func(r *worker.Result) {
		fmt.Printf("Result: %d, Error: %v\n", r.Value, r.Error)
	}

	w.FinishedHandler = func() {
		fmt.Println("finished task queue")
	}

	w.Start(25)
}
