package e

import "fmt"

type Error struct {
        msg  string
}

func (e *Error) Error() string {
    return e.msg
}

func Wrap(msg string, err error) error{
	return  fmt.Errorf("%s: %w",msg, err)
}

func WrapIfErr(msg string, err error) error{
	if err == nil {
		return nil
	}
	return  Wrap(msg,err)
}

func New(msg string) *Error {
    return &Error{msg: msg}
}