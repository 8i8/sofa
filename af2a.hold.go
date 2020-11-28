package sofa

import (
	"errors"
	"math"
)

var (
	errAf2aOne   = errors.New("ideg outside range 0-359      ")
	errAf2aTwo   = errors.New("iamin outside range 0-59      ")
	errAf2aThree = errors.New("asec outside range 0-59.999...")
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
// int iauAf2a(char s, int ideg, int iamin, double asec, double *rad)
func Af2a(s byte, ideg, iamin int, asec float64) (rad float64, err error) {
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
		err = errAf2aOne
	}
	if iamin < 0 || iamin > 59 {
		err = errAf2aTwo
	}
	if asec < 0.0 || asec >= 60.0 {
		err = errAf2aThree
	}
	return
}
