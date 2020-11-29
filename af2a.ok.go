package sofa

// #include <stdio.h>
// #include "sofa.h"
import "C"
import (
	"errors"
	"math"
)

var (
	errAf2aIdeg  = errors.New("ideg outside range 0-359")
	errAf2aIamin = errors.New("iamin outside range 0-59")
	errAf2aAsec  = errors.New("asec outside range 0-59.999...")
)

//  Af2a Convert degrees, arcminutes, arcseconds to radians.
//
//  - - - - -
//   A f 2 a
//  - - - - -
//
//  This function is part of the International Astronomical Union's
//  SOFA (Standards of Fundamental Astronomy) software collection.
//
//  Status:  support function.
//
//  Given:
//     s         char    sign:  '-' = negative, otherwise positive
//     ideg      int     degrees
//     iamin     int     arcminutes
//     asec      double  arcseconds
//
//  Returned:
//     rad       double  angle in radians
//
//  Returned (function value):
//               int     status:  0 = OK
//                                1 = ideg outside range 0-359
//                                2 = iamin outside range 0-59
//                                3 = asec outside range 0-59.999...
//
//  Notes:
//
//  1)  The result is computed even if any of the range checks fail.
//
//  2)  Negative ideg, iamin and/or asec produce a warning status, but
//      the absolute value is used in the conversion.
//
//  3)  If there are multiple errors, the status value reflects only the
//      first, the smallest taking precedence.
//
//  This revision:  2013 June 18
//
//  SOFA release 2020-07-21
//
//  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
//
func Af2a(s byte, ideg, iamin int, asec float64) (rad float64, err error) {
	var cRad C.double
	j := C.iauAf2a(C.char(s), C.int(ideg), C.int(iamin), C.double(asec), &cRad)
	switch int(j) {
	case 1:
		err = errAf2aIdeg
	case 2:
		err = errAf2aIamin
	case 3:
		err = errAf2aAsec
	}
	return float64(cRad), err
}

func goAf2a(s byte, ideg, iamin int, asec float64) (rad float64, err error) {
	/* Compute the interval. */
	var sign = 1.0
	if s == '-' {
		sign = -sign
	}

	rad = sign * (60.0*(60.0*
		(float64(Abs(ideg)))+(float64(Abs(iamin)))) +
		math.Abs(asec)) * DAS2R

	/* Validate arguments and return status. */
	if ideg < 0 || ideg > 359 {
		err = errAf2aIdeg
	}
	if iamin < 0 || iamin > 59 {
		err = errAf2aIamin
	}
	if asec < 0.0 || asec >= 60.0 {
		err = errAf2aAsec
	}
	return
}
