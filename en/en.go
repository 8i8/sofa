package en

import (
	"bytes"
	"log"
	"os"
)

// ErrNum is an error type for dealing with enumerated errors.
type ErrNum interface {
	Error() string
	Set(int) ErrNum
	Is() int
	Add(ErrNum, int) ErrNum
	Wrap(ErrNum) ErrNum
	Name() string
}

type errnum struct {
	n      int
	offset int // zero offset, to allow for negative array indices.
	name   string
	msg    []string
}

// Set sets the error value that is to be output to the user.
func (e errnum) Set(n int) ErrNum {
	e.check(n)
	e.n = n
	return e
}

// Add additions v with the current error value.
func (e errnum) Add(err ErrNum, v int) ErrNum {
	if err == nil {
		e.check(v)
		e = e.Set(v).(errnum)
		return e
	}
	e.n += (err.Is() + v)
	e.check(e.n)
	return e
}

// GetNum returns the current set error value.
func (e errnum) GetNum() int {
	return e.n
}

// Is returns the errors value.
func (e errnum) Is() int {
	return e.n
}

// Wrap mantains an errrors value and adds the previous errors name to
// the current name, replacing the mesage with that if the current
// error.
func (e errnum) Wrap(err ErrNum) ErrNum {
	e.name = err.Name() + ": " + e.name
	return e
}

// Name returns the name of an ErrNum error, generaly the name of the
// function that the error is ascribed to.
func (e errnum) Name() string {
	return e.name
}

// Error returns the appropriate error message.
func (e errnum) Error() string {
	buf := bytes.Buffer{}
	if len(e.name) > 0 {
		buf.WriteString(e.name)
	}
	switch e.status() {
	case 2:
		buf.WriteString(" err.Error(): value of 'n' too high: ")
	case 1:
		buf.WriteString(" warning: ")
	case 0:
		buf.WriteString(" please contact package administration: ")
	case -1:
		buf.WriteString(" error: ")
	case -2:
		buf.WriteString(" err.Error(): value of 'n' too low: ")
	}
	buf.WriteString(e.msg[e.n+e.offset])
	return buf.String()
}

// status returns 1 to indicate any valid posative state as a warning
// and -1 to indicate a valid error. 2 is returned if the status is
// erronious due to the number being too great, and -2 if it is too low.
// a state of 0 is returnd if the error state is 0.
func (e errnum) status() int {
	switch {
	case e.n > len(e.msg)-e.offset-1:
		// e.n is greater than the highest output index.
		return 2
	case e.n > 0:
		// e.n is posative.
		return 1
	case e.n < -e.offset:
		// e.n is lower than the least index.
		return -2
	case e.n < 0:
		// e.n is negative.
		return -1
	default:
		// unknown value, error.
		return 0
	}
}

// New returns a new ErrNum, setting output messages, value and the
// offset from 0 for dealing with negative value errors.
func New(o int, name string, msg []string) ErrNum {
	if o > len(msg) {
		log.SetFlags(log.Lshortfile)
		log.Output(3, "invalid offset")
		os.Exit(1)
	}
	return errnum{offset: o, name: name, msg: msg}
}

// check verifies that the current error value is within the output
// boundaries, if it is not then a fatal error is triggered.
func (e errnum) check(v int) {
	if v >= len(e.msg)-e.offset {
		log.SetFlags(log.Lshortfile)
		log.Output(3, "value 'Set(n)' to high")
		os.Exit(1)
	}
	if v < -e.offset {
		log.SetFlags(log.Lshortfile)
		log.Output(3, "value 'Set(n)' to low")
		os.Exit(1)
	}
}
