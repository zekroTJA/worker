package worker

// Result contains the return value of the
// tasks function as interface and an error
// instance if the tasks function failes.
type Result struct {
	Value interface{}
	Error error
}

// NewResult builds a new instance of Result
// with the passed value and error instance.
func NewResult(value interface{}, err error) *Result {
	return &Result{
		Value: value,
		Error: err,
	}
}
