package helper

import (
	"fmt"
)

type Err struct {
	Func string
	Path string
	Err  error
}

func (e *Err) Error() string {
	return e.ErrorSteps(3)
}

// Use -1 to see all
func (e *Err) ErrorSteps(last int) string {
	list, final := MyErrList(e)
	if last > 0 && last < len(list) {
		list = list[len(list)-last:]
	}
	steps := ""
	for _, myErr := range list {
		steps += fmt.Sprintf(
			"fn:  %v%vpath: %v\n", myErr.Func, spaces(myErr.Func), myErr.Path,
		)
	}
	return fmt.Sprintf("%vERR: %v\n", steps, final)
}

func MyErrList(current error) ([]*Err, error) {
	myErrs := []*Err{}

	var err error
	var myErr *Err

	for {
		myErr, err = ErrCheck(current)

		if err != nil || (myErr == nil && err == nil) {
			break
		}

		current = myErr.Err
		myErrs = append(myErrs, myErr)
	}

	return myErrs, err
}

func ErrCheck(e error) (*Err, error) {
	myErr, ok := e.(*Err)
	if ok {
		return myErr, nil
	}
	return nil, e
}

func spaces(fn string) string {
	n := 18 - len(fn)
	str := ""
	for i := 0; i <= n; i++ {
		str += " "
	}
	return str
}
