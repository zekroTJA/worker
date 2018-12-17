<div align="center">
    <h1>~ worker ~</h1>
    <strong>Thread pool package for go routines</strong><br><br>
    <a href="https://godoc.org/github.com/zekroTJA/worker"><img src="https://img.shields.io/badge/docs-godoc-c918cc.svg" /></a>&nbsp;
    <a href="https://travis-ci.org/zekroTJA/worker" ><img src="https://travis-ci.org/zekroTJA/worker.svg?branch=master" /></a>&nbsp;
    <a href="https://coveralls.io/github/zekroTJA/worker"><img src="https://coveralls.io/repos/github/zekroTJA/worker/badge.svg" /></a>
<br>
</div>

---

<div align="center">
    <code>go get github.com/zekroTJA/worker</code>
</div>

---

## Intro

This is a leight go package to split work load into small tasks, which can be spread over a pool of workers which will proceed the tasks in different, concurrent go routines.

[Here](https://godoc.org/github.com/zekroTJA/worker) you can read the docs of this package, generated by godoc.org.

---

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/zekroTJA/worker"
)

func main() {
    // Create worker instance
    w := worker.NewWorker()

    // Define 50 Tasks, which should be proceed, and 
    // enqueue them in the workers queue
    for j := 0; j < 50; j++ {
        task := worker.NewTask(func(args []interface{}) *worker.Result {
            time.Sleep(1 * time.Second)
            return worker.NewResult(args[0], nil)
        }
        w.Enqueue(task, j))
    }

    // Define the result handler which will be executed each time
    // a worker has finished a task and returns a result instance
    w.ResultHandler = func(r *worker.Result) {
        fmt.Printf("Result: %d, Error: %v\n", r.Value, r.Error)
    }

    // Define the finish handler which will be executed after all
    // tasks in the queue are finished
    w.FinishedHandler = func() {
        fmt.Println("finished task queue")
    }

    // Start processing the queue with a pool of 25 workers, which will
    // works simultaniously on processing the queue.
    // This function will block the current thread until all tasks in the
    // queue are finished.
    w.Start(25)
}
```

Further examples, you can find in the [example](example) directory.

---

Copyright (c) 2018 zekro Development (Ringo Hoffmann).
Covered with MIT licence.
