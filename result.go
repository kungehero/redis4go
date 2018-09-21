package redis4go

type Result struct {
	Data interface{}
	err  error
}

func result(data interface{}, err error) *Result {
	return &Result{data, err}
}
