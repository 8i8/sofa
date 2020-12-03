package sofa

import "errors"

var (
	errAdmin = errors.New("something went wrong, " +
		"please inform the package maintainer")
)
