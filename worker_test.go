package worker

import "testing"

var worker *Worker

var results = make(map[int]int)

func TestNewWorker(t *testing.T) {
	worker = NewWorker(100)
	if worker == nil {
		t.Fatal("created worker was nil")
	}
}

func TestEnqueue(t *testing.T) {
	for i := 1; i <= 20; i++ {
		task := NewTask(func(args []interface{}) *Result {
			return NewResult([]int{args[0].(int), args[0].(int) * 2}, nil)
		}, i)
		worker.Enqueue(task)
	}
}

func TestSetResultHandler(t *testing.T) {
	worker.SetResultHandler(func(r *Result) {
		v := r.Value.([]int)
		results[v[0]] = v[1]
	})
}

func TestSetFinishedHandler(t *testing.T) {
	worker.SetFinishedHandler(func() {
		for i := 1; i <= 20; i++ {
			if v, ok := results[i]; !ok {
				t.Fatalf("key %d was not found in results map", i)
			} else if v != i*2 {
				t.Fatalf("value was %d and should be %d", v, i*2)
			}
		}
	})
}

func TestStart(t *testing.T) {
	worker.Start(5)
}
