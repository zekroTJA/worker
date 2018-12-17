package worker

// Task contains the function to execute, which is
// getting passed arguments as array of interfaces and
// returns a result instance and the actual array of
// arguments, which are getting passed to the tasks
// function.
type Task struct {
	F    func([]interface{}) *Result
	Args []interface{}
}

// NewTask creates a new Task instance with the function
// to proceed and the arguments, which will be passed to
// the tasks function.
func NewTask(f func([]interface{}) *Result, args ...interface{}) *Task {
	return &Task{
		F:    f,
		Args: args,
	}
}
