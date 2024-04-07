package wrap

type ErrorWrap struct {
	Msg string
	Err error
}

func (e *ErrorWrap) Error() string {
	return e.Msg
}

func (e *ErrorWrap) Unwrap() error {
	return e.Err
}
