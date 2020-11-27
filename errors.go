package sofa

import "errors"

var (
	ErrYear    = errors.New("bad year")
	ErrMonth   = errors.New("bad month")
	ErrDay     = errors.New("bad day")
	ErrWarning = errors.New("warning")
)
